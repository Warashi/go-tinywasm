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

	"github.com/Warashi/wasmium/runtime"
	typesRuntime "github.com/Warashi/wasmium/types/runtime"
)

type JSONWast struct {
	SourceFilename string     `json:"source_filename"`
	Commands       []Commands `json:"commands"`
}
type Args struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Action struct {
	Type  string `json:"type"`
	Field string `json:"field"`
	Args  []Args `json:"args"`
}
type Expected struct {
	Type  string `json:"type"`
	Value string `json:"value"`
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

type Result = Expected

func action(r *runtime.Runtime, a Action) ([]Result, error) {
	switch a.Type {
	case "invoke":
		return invoke(r, a)
	}
	return nil, fmt.Errorf("unsupported action type %s", a.Type)
}

// convert converts from F to T with same binary representation.
func convert[T, F any](f F, bit int) T {
	buf := make([]byte, bit/8)
	binary.Encode(buf, binary.LittleEndian, f)

	var t T
	binary.Decode(buf, binary.LittleEndian, &t)
	return t
}

func invoke(r *runtime.Runtime, a Action) ([]Result, error) {
	field := a.Field
	args := make([]typesRuntime.Value, len(a.Args))

	for i, arg := range a.Args {
		switch arg.Type {
		case "i32":
			a, err := strconv.ParseInt(arg.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			args[i] = typesRuntime.ValueI32(a)
		case "i64":
			a, err := strconv.ParseInt(arg.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			args[i] = typesRuntime.ValueI64(a)
		case "f32":
			a, err := strconv.ParseFloat(arg.Value, 32)
			if err != nil {
				return nil, err
			}
			args[i] = typesRuntime.ValueF32(a)
		case "f64":
			a, err := strconv.ParseFloat(arg.Value, 64)
			if err != nil {
				return nil, err
			}
			args[i] = typesRuntime.ValueF64(a)
		}
	}

	v, err := r.Call(field, args...)
	if err != nil {
		return nil, err
	}

	var result []Result
	for _, v := range v {
		switch v := v.(type) {
		case typesRuntime.ValueI32:
			result = append(result, Result{Type: "i32", Value: strconv.FormatInt(int64(v), 10)})
		case typesRuntime.ValueI64:
			result = append(result, Result{Type: "i64", Value: strconv.FormatInt(int64(v), 10)})
		case typesRuntime.ValueF32:
			result = append(result, Result{Type: "f32", Value: strconv.FormatUint(uint64(convert[uint32](v, 32)), 10)})
		case typesRuntime.ValueF64:
			result = append(result, Result{Type: "f64", Value: strconv.FormatUint(convert[uint64](v, 64), 10)})
		}
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
	baseDir := cmp.Or(os.Getenv("WASMIUM_TEST_DIR"), ".")

	for p := range filepathWalk(t, baseDir) {
		if filepath.Ext(p) != ".json" {
			continue
		}

		t.Run(p, func(t *testing.T) {
			wast := setup(t, p)

			var r *runtime.Runtime

			for _, cmd := range wast.Commands {
				switch cmd.Type {
				case "module":
					f, err := os.Open(filepath.Join(baseDir, cmd.Filename))
					if err != nil {
						t.Fatalf("failed to open file %s: %v", cmd.Filename, err)
					}
					r, err = runtime.New(f)
					if err != nil {
						t.Fatalf("failed to create runtime: %v", err)
					}
				case "assert_return":
					got, err := action(r, cmd.Action)
					if err != nil {
						t.Errorf("failed to execute action: %v", err)
					}
					if !reflect.DeepEqual(cmd.Expected, got) {
						t.Errorf("assertion failed: expected %v, got %v", cmd.Expected, got)
					}
				}
			}
		})
	}
}
