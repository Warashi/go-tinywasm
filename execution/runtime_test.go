package execution_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/Warashi/go-tinywasm/execution"
)

func TestExecuteI32Add(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_add.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	runtime, err := execution.NewRuntime(bytes.NewReader(b))
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
			args := []execution.Value{
				execution.ValueI32(test.a),
				execution.ValueI32(test.b),
			}
			got, err := runtime.Call(0, args)
			if err != nil {
				t.Errorf("failed to call function: %v", err)
				t.FailNow()
			}
			if len(got) != 1 {
				t.Errorf("unexpected number of return values: %d", len(got))
				t.FailNow()
			}
			if got, ok := got[0].(execution.ValueI32); !ok || got != execution.ValueI32(test.want) {
				t.Errorf("unexpected return value: %v", got)
			}
		})
	}
}
