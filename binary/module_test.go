package binary

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/Warashi/wasmium/instruction"
	"github.com/Warashi/wasmium/types/binary"
)

func TestDecodePreamble(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/minimal.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	m, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	if m.magic != "\x00asm" {
		t.Errorf("wrong magic bytes: %x", m.magic)
	}
	if m.version != 1 {
		t.Errorf("wrong version: %d", m.version)
	}
}

func TestDecodeMinimalFunc(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/minimal_func.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	got, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	want := &Module{
		magic:   "\x00asm",
		version: 1,
		typeSection: []binary.FuncType{
			{
				Params:  []binary.ValueType{},
				Results: []binary.ValueType{},
			},
		},
		functionSection: []uint32{0},
		codeSection: []binary.Function{
			{
				Locals: []binary.FunctionLocal{},
				Code:   []binary.Instruction{&instruction.End{}},
			},
		},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected Module: %#v", got)
	}
}

func TestDecodeFuncParam(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_param.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	got, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	want := &Module{
		magic:   "\x00asm",
		version: 1,
		typeSection: []binary.FuncType{
			{
				Params:  []binary.ValueType{binary.ValueTypeI32, binary.ValueTypeI64},
				Results: []binary.ValueType{},
			},
		},
		functionSection: []uint32{0},
		codeSection: []binary.Function{
			{
				Locals: []binary.FunctionLocal{},
				Code:   []binary.Instruction{&instruction.End{}},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected Module: %#v", got)
	}
}

func TestDecodeFuncLocal(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_local.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	got, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	want := &Module{
		magic:   "\x00asm",
		version: 1,
		typeSection: []binary.FuncType{
			{
				Params:  []binary.ValueType{},
				Results: []binary.ValueType{},
			},
		},
		functionSection: []uint32{0},
		codeSection: []binary.Function{
			{
				Locals: []binary.FunctionLocal{
					{TypeCount: 1, ValueType: binary.ValueTypeI32},
					{TypeCount: 2, ValueType: binary.ValueTypeI64},
				},
				Code: []binary.Instruction{&instruction.End{}},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected Module: %#v", got)
	}
}

func TestDecodeFuncAdd(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_add.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	got, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	want := &Module{
		magic:   "\x00asm",
		version: 1,
		typeSection: []binary.FuncType{
			{
				Params:  []binary.ValueType{binary.ValueTypeI32, binary.ValueTypeI32},
				Results: []binary.ValueType{binary.ValueTypeI32},
			},
		},
		functionSection: []uint32{0},
		codeSection: []binary.Function{
			{
				Locals: []binary.FunctionLocal{},
				Code: []binary.Instruction{
					&instruction.LocalGet{Index: 0},
					&instruction.LocalGet{Index: 1},
					&instruction.I32Add{},
					&instruction.End{},
				},
			},
		},
		exportSection: []binary.Export{
			{
				Name: "add",
				Desc: binary.ExportDescFunc{Index: 0},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected Module: %#v", got)
	}
}

func TestDecodeFuncCall(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/func_call.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	got, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	want := &Module{
		magic:   "\x00asm",
		version: 1,
		typeSection: []binary.FuncType{
			{
				Params:  []binary.ValueType{binary.ValueTypeI32},
				Results: []binary.ValueType{binary.ValueTypeI32},
			},
		},
		functionSection: []uint32{0, 0},
		codeSection: []binary.Function{
			{
				Locals: []binary.FunctionLocal{},
				Code: []binary.Instruction{
					&instruction.LocalGet{Index: 0},
					&instruction.Call{Index: 1},
					&instruction.End{},
				},
			},
			{
				Locals: []binary.FunctionLocal{},
				Code: []binary.Instruction{
					&instruction.LocalGet{Index: 0},
					&instruction.LocalGet{Index: 0},
					&instruction.I32Add{},
					&instruction.End{},
				},
			},
		},
		exportSection: []binary.Export{
			{
				Name: "call_doubler",
				Desc: binary.ExportDescFunc{Index: 0},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected Module: %#v", got)
	}
}

func TestDecodeImport(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/import.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	got, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	want := &Module{
		magic:   "\x00asm",
		version: 1,
		typeSection: []binary.FuncType{
			{
				Params:  []binary.ValueType{binary.ValueTypeI32},
				Results: []binary.ValueType{binary.ValueTypeI32},
			},
		},
		importSection: []binary.Import{
			{
				Module: "env",
				Field:  "add",
				Desc:   binary.ImportDescFunc{Index: 0},
			},
		},
		exportSection: []binary.Export{
			{
				Name: "call_add",
				Desc: binary.ExportDescFunc{Index: 1},
			},
		},
		functionSection: []uint32{0},
		codeSection: []binary.Function{
			{
				Locals: []binary.FunctionLocal{},
				Code: []binary.Instruction{
					&instruction.LocalGet{Index: 0},
					&instruction.Call{Index: 0},
					&instruction.End{},
				},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected Module: %#v", got)
	}
}

func TestDecodeFib(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("../testdata/fib.wasm")
	if err != nil {
		t.Errorf("failed to load testdata: %v", err)
		t.FailNow()
	}

	got, err := NewModule(bytes.NewReader(b))
	if err != nil {
		t.Errorf("failed to parse wasm: %v", err)
	}

	want := &Module{
		magic:   "\x00asm",
		version: 1,
		typeSection: []binary.FuncType{
			{
				Params:  []binary.ValueType{binary.ValueTypeI32},
				Results: []binary.ValueType{binary.ValueTypeI32},
			},
		},
		exportSection: []binary.Export{
			{
				Name: "fib",
				Desc: binary.ExportDescFunc{Index: 0},
			},
		},
		functionSection: []uint32{0},
		codeSection: []binary.Function{
			{
				Locals: []binary.FunctionLocal{},
				Code: []binary.Instruction{
					&instruction.LocalGet{Index: 0},
					&instruction.I32Const{Value: 2},
					&instruction.I32LtS{},
					&instruction.If{Block: binary.Block{BlockType: binary.BlockTypeVoid{}}},
					&instruction.I32Const{Value: 1},
					&instruction.Return{},
					&instruction.End{},
					&instruction.LocalGet{Index: 0},
					&instruction.I32Const{Value: 2},
					&instruction.I32Sub{},
					&instruction.Call{Index: 0},
					&instruction.LocalGet{Index: 0},
					&instruction.I32Const{Value: 1},
					&instruction.I32Sub{},
					&instruction.Call{Index: 0},
					&instruction.I32Add{},
					&instruction.Return{},
					&instruction.End{},
				},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected Module: %#v", got)
	}
}
