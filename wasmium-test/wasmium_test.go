package wasmium_test

import (
	"cmp"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"
	"unsafe"

	"github.com/Warashi/wasmium/runtime"
	typesRuntime "github.com/Warashi/wasmium/types/runtime"
)

type JSONWast struct {
	SourceFilename string     `json:"source_filename"`
	Commands       []Commands `json:"commands"`
}
type Value struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

func (a Value) RuntimeValue() typesRuntime.Value {
	switch a.Type {
	case "i32":
		return typesRuntime.ValueI32(a.i32())
	case "i64":
		return typesRuntime.ValueI64(a.i64())
	case "f32":
		return typesRuntime.ValueF32(a.f32())
	case "f64":
		return typesRuntime.ValueF64(a.f64())
	}
	panic(fmt.Sprintf("unsupported value type %s", a.Type))
}

func (a Value) i32() int32 {
	var s string
	json.Unmarshal(a.Value, &s)
	v, _ := strconv.ParseInt(s, 10, 32)
	return int32(v)
}

func (a Value) i64() int64 {
	var s string
	json.Unmarshal(a.Value, &s)
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func (a Value) f32() float32 {
	var s string
	json.Unmarshal(a.Value, &s)
	v, _ := strconv.ParseUint(s, 10, 32)
	return convert[float32](v)
}

func (a Value) f64() float64 {
	var s string
	json.Unmarshal(a.Value, &s)
	v, _ := strconv.ParseUint(s, 10, 64)
	return convert[float64](v)
}

func (a Value) String() string {
	switch a.Type {
	case "i32":
		var buf [4]byte
		binary.Encode(buf[:], binary.LittleEndian, a.i32())
		return fmt.Sprintf("%s(0x%x)", a.Type, buf)
	case "i64":
		var buf [8]byte
		binary.Encode(buf[:], binary.LittleEndian, a.i64())
		return fmt.Sprintf("%s(0x%x)", a.Type, buf)
	case "f32":
		var buf [4]byte
		binary.Encode(buf[:], binary.LittleEndian, a.i32())
		return fmt.Sprintf("%s(0x%x)", a.Type, buf)
	case "f64":
		var buf [8]byte
		binary.Encode(buf[:], binary.LittleEndian, a.i64())
		return fmt.Sprintf("%s(0x%x)", a.Type, buf)
	}
	panic(fmt.Sprintf("unsupported value type %s", a.Type))
}

func mustMarshalJSON(v interface{}) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal json: %v", err))
	}
	return json.RawMessage(b)
}

func NewValue(v typesRuntime.Value) Value {
	switch v := v.(type) {
	case typesRuntime.ValueI32:
		return Value{Type: "i32", Value: mustMarshalJSON(strconv.FormatInt(int64(v), 10))}
	case typesRuntime.ValueI64:
		return Value{Type: "i64", Value: mustMarshalJSON(strconv.FormatInt(int64(v), 10))}
	case typesRuntime.ValueF32:
		return Value{Type: "f32", Value: mustMarshalJSON(strconv.FormatUint(uint64(convert[uint32](v)), 10))}
	case typesRuntime.ValueF64:
		return Value{Type: "f64", Value: mustMarshalJSON(strconv.FormatUint(convert[uint64](v), 10))}
	default:
		panic(fmt.Sprintf("unsupported value type %T", v))
	}
}

type Action struct {
	Type  string  `json:"type"`
	Field string  `json:"field"`
	Args  []Value `json:"args"`
}

func (a Action) String() string {
	return fmt.Sprintf("%s %s %v", a.Type, a.Field, a.Args)
}

type Commands struct {
	Type       string     `json:"type"`
	Line       int        `json:"line"`
	Filename   string     `json:"filename,omitempty"`
	Action     Action     `json:"action,omitempty"`
	Expected   []Expected `json:"expected,omitempty"`
	Text       string     `json:"text,omitempty"`
	ModuleType string     `json:"module_type,omitempty"`
}

func (c Commands) TestName() string {
	return fmt.Sprintf("%s:%d(%s:%s)", c.Filename, c.Line, c.Type, c.Action)
}

type Expected = Value
type Result = Expected

func action(r *runtime.Runtime, a Action) ([]Result, error) {
	switch a.Type {
	case "invoke":
		return invoke(r, a)
	}
	return nil, fmt.Errorf("unsupported action type %s", a.Type)
}

// convert converts from F to T with same binary representation.
func convert[T, F any](f F) T {
	var t T
	if unsafe.Sizeof(f) != unsafe.Sizeof(t) {
		panic(fmt.Sprintf("size mismatch: %d != %d", unsafe.Sizeof(f), unsafe.Sizeof(t)))
	}
	return *(*T)(unsafe.Pointer(&f))
}

func invoke(r *runtime.Runtime, a Action) ([]Result, error) {
	field := a.Field
	args := make([]typesRuntime.Value, len(a.Args))

	for i, arg := range a.Args {
		args[i] = arg.RuntimeValue()
	}

	v, err := r.Call(field, args...)
	if err != nil {
		return nil, err
	}

	var result []Result
	for _, v := range v {
		result = append(result, NewValue(v))
	}

	return result, nil
}

func filepathWalk(t *testing.T, basedir string) func(func(string) bool) {
	t.Helper()

	return func(yield func(string) bool) {
		filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				t.Fatalf("failed to walk %s: %v", path, err)
				return err
			}
			if info.IsDir() {
				return nil
			}

			if !yield(path) {
				return filepath.SkipAll
			}
			return nil
		})
	}
}

func setup(t *testing.T, p string) JSONWast {
	t.Helper()

	f, err := os.Open(p)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", p, err)
	}
	defer f.Close()

	var wast JSONWast
	if err := json.NewDecoder(f).Decode(&wast); err != nil {
		t.Fatalf("failed to decode json: %v", err)
	}

	return wast
}

func TestWasmium(t *testing.T) {
	t.Parallel()

	baseDir := cmp.Or(os.Getenv("WASMIUM_TEST_DIR"), ".")

	for p := range filepathWalk(t, baseDir) {
		if filepath.Ext(p) != ".json" {
			continue
		}

		t.Run(filepath.Base(p), func(t *testing.T) {
			t.Parallel()

			wast := setup(t, p)

			var r *runtime.Runtime

			for _, cmd := range wast.Commands {
				t.Run(cmd.TestName(), func(t *testing.T) {
					switch cmd.Type {
					case "module":
						f, err := os.Open(filepath.Join(baseDir, cmd.Filename))
						if err != nil {
							t.Fatalf("failed to open file %s: %v", cmd.Filename, err)
						}
						defer f.Close()

						r, err = runtime.New(f)
						if err != nil {
							t.Fatalf("failed to create runtime: %v", err)
						}
					case "assert_return":
						if r == nil {
							t.Skip("module loading failed")
						}
						got, err := action(r, cmd.Action)
						if err != nil {
							t.Errorf("failed to execute action: %v", err)
						}
						if !reflect.DeepEqual(cmd.Expected, got) {
							t.Errorf("assertion failed: expected %v, got %v", cmd.Expected, got)
						}
					}
				})
			}
		})
	}
}
