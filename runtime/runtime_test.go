package runtime_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/Warashi/go-tinywasm/runtime"

	typesRuntime "github.com/Warashi/go-tinywasm/types/runtime"
)

func TestExecuteI32Add(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_add.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	tests := []struct {
		a, b int32
		want int32
	}{
		{1, 2, 3},
		{3, 4, 7},
		{5, 6, 11},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d+%d=%d", test.a, test.b, test.want), func(t *testing.T) {
			args := []typesRuntime.Value{
				typesRuntime.ValueI32(test.a),
				typesRuntime.ValueI32(test.b),
			}
			got, err := runtime.Call("add", args...)
			if err != nil {
				t.Errorf("failed to call function: %v", err)
				t.FailNow()
			}
			if len(got) != 1 {
				t.Errorf("unexpected number of return values: %d", len(got))
				t.FailNow()
			}
			if got, ok := got[0].(typesRuntime.ValueI32); !ok || got != typesRuntime.ValueI32(test.want) {
				t.Errorf("unexpected return value: %v", got)
			}
		})
	}
}

func TestNotFoundExportedFunction(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_add.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	if _, err := runtime.Call("not_found", nil); err == nil {
		t.Error("unexpected success")
	}
}

func TestFuncCall(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_call.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	tests := []struct {
		a    int32
		want int32
	}{
		{1, 2},
		{2, 4},
		{3, 6},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("double(%d)=%d", test.a, test.want), func(t *testing.T) {
			args := []typesRuntime.Value{
				typesRuntime.ValueI32(test.a),
			}
			got, err := runtime.Call("call_doubler", args...)
			if err != nil {
				t.Errorf("failed to call function: %v", err)
				t.FailNow()
			}
			if len(got) != 1 {
				t.Errorf("unexpected number of return values: %d", len(got))
				t.FailNow()
			}
			if got, ok := got[0].(typesRuntime.ValueI32); !ok || got != typesRuntime.ValueI32(test.want) {
				t.Errorf("unexpected return value: %v", got)
			}
		})
	}
}

func TestCallImportedFunc(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/import.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	r, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	r.AddImport("env", "add", func(s *runtime.Store, v ...typesRuntime.Value) ([]typesRuntime.Value, error) {
		switch arg := v[0].(type) {
		case typesRuntime.ValueI32:
			return []typesRuntime.Value{typesRuntime.ValueI32(arg + arg)}, nil
		default:
			return nil, fmt.Errorf("unsupported argument type: %T", arg)
		}
	})

	tests := []struct {
		a, want int32
	}{
		{1, 2},
		{2, 4},
		{3, 6},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("double(%d)=%d", test.a, test.want), func(t *testing.T) {
			args := []typesRuntime.Value{
				typesRuntime.ValueI32(test.a),
			}
			got, err := r.Call("call_add", args...)
			if err != nil {
				t.Errorf("failed to call function: %v", err)
				t.FailNow()
			}
			if len(got) != 1 {
				t.Errorf("unexpected number of return values: %d", len(got))
				t.FailNow()
			}
			if got, ok := got[0].(typesRuntime.ValueI32); !ok || got != typesRuntime.ValueI32(test.want) {
				t.Errorf("unexpected return value: %v", got)
			}
		})
	}
}

func TestNotFoundImportedFunc(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/import.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	if _, err := runtime.Call("call_not_found", nil); err == nil {
		t.Error("unexpected success")
	}
}

func TestI32Const(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/i32_const.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	got, err := runtime.Call("i32_const", nil)
	if err != nil {
		t.Errorf("failed to call function: %v", err)
		t.FailNow()
	}

	if len(got) != 1 {
		t.Errorf("unexpected number of return values: %d", len(got))
		t.FailNow()
	}

	if got, ok := got[0].(typesRuntime.ValueI32); !ok || got != 42 {
		t.Errorf("unexpected return value: %v", got)
	}
}

func TestLocalSet(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/local_set.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	got, err := runtime.Call("local_set", nil)
	if err != nil {
		t.Errorf("failed to call function: %v", err)
		t.FailNow()
	}

	if len(got) != 1 {
		t.Errorf("unexpected number of return values: %d", len(got))
		t.FailNow()
	}

	if got, ok := got[0].(typesRuntime.ValueI32); !ok || got != 42 {
		t.Errorf("unexpected return value: %v", got)
	}
}

func TestI32Store(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/i32_store.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	if _, err := runtime.Call("i32_store", nil); err != nil {
		t.Errorf("failed to call function: %v", err)
		t.FailNow()
	}

	memory := runtime.Store().Memories()[0].Data
	if memory[0] != 42 {
		t.Errorf("unexpected memory content: %v", memory)
	}
}

func TestI32Sub(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_sub.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	tests := []struct {
		a, b int32
		want int32
	}{
		{1, 2, -1},
		{3, 4, -1},
		{5, 6, -1},
		{10, 5, 5},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d-%d=%d", test.a, test.b, test.want), func(t *testing.T) {
			args := []typesRuntime.Value{
				typesRuntime.ValueI32(test.a),
				typesRuntime.ValueI32(test.b),
			}
			got, err := runtime.Call("sub", args...)
			if err != nil {
				t.Errorf("failed to call function: %v", err)
				t.FailNow()
			}
			if len(got) != 1 {
				t.Errorf("unexpected number of return values: %d", len(got))
				t.FailNow()
			}
			if got, ok := got[0].(typesRuntime.ValueI32); !ok || got != typesRuntime.ValueI32(test.want) {
				t.Errorf("unexpected return value: %v", got)
			}
		})
	}
}

func TestI32LessThan(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_lts.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	tests := []struct {
		a, b int32
		want int32
	}{
		{1, 2, 1},
		{3, 4, 1},
		{6, 5, 0},
		{10, 10, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d<%d=%d", test.a, test.b, test.want), func(t *testing.T) {
			args := []typesRuntime.Value{
				typesRuntime.ValueI32(test.a),
				typesRuntime.ValueI32(test.b),
			}
			got, err := runtime.Call("lts", args...)
			if err != nil {
				t.Errorf("failed to call function: %v", err)
				t.FailNow()
			}
			if len(got) != 1 {
				t.Errorf("unexpected number of return values: %d", len(got))
				t.FailNow()
			}
			if got, ok := got[0].(typesRuntime.ValueI32); !ok || got != typesRuntime.ValueI32(test.want) {
				t.Errorf("unexpected return value: %v", got)
			}
		})
	}
}

func TestFib(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/fib.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := runtime.New(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to create runtime: %v", err)
		t.FailNow()
	}

	tests := []struct {
		n, want int32
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{7, 21},
		{8, 34},
		{9, 55},
		{10, 89},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("fib(%d)=%d", test.n, test.want), func(t *testing.T) {
			args := []typesRuntime.Value{
				typesRuntime.ValueI32(test.n),
			}
			got, err := runtime.Call("fib", args...)
			if err != nil {
				t.Errorf("failed to call function: %v", err)
				t.FailNow()
			}
			if len(got) != 1 {
				t.Errorf("unexpected number of return values: %d", len(got))
				t.FailNow()
			}
			if got, ok := got[0].(typesRuntime.ValueI32); !ok || got != typesRuntime.ValueI32(test.want) {
				t.Errorf("unexpected return value: %v", got)
			}
		})
	}
}
