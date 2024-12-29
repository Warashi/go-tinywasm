package binary

import (
	"bytes"
	"os"
	"reflect"
	"testing"
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
		typeSection: []FuncType{
			{
				params:  []ValueType{},
				results: []ValueType{},
			},
		},
		functionSection: []uint32{0},
		codeSection: []Function{
			{
				locals: []FunctionLocal{},
				code:   []Instruction{InstructionEnd{}},
			},
		},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected module: %#v", got)
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
		typeSection: []FuncType{
			{
				params:  []ValueType{ValueTypeI32, ValueTypeI64},
				results: []ValueType{},
			},
		},
		functionSection: []uint32{0},
		codeSection: []Function{
			{
				locals: []FunctionLocal{},
				code:   []Instruction{InstructionEnd{}},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected module: %#v", got)
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
		typeSection: []FuncType{
			{
				params:  []ValueType{},
				results: []ValueType{},
			},
		},
		functionSection: []uint32{0},
		codeSection: []Function{
			{
				locals: []FunctionLocal{
					{typeCount: 1, valueType: ValueTypeI32},
					{typeCount: 2, valueType: ValueTypeI64},
				},
				code: []Instruction{InstructionEnd{}},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected module: %#v", got)
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
		typeSection: []FuncType{
			{
				params:  []ValueType{ValueTypeI32, ValueTypeI32},
				results: []ValueType{ValueTypeI32},
			},
		},
		functionSection: []uint32{0},
		codeSection: []Function{
			{
				locals: []FunctionLocal{},
				code: []Instruction{
					&InstructionLocalGet{idx: 0},
					&InstructionLocalGet{idx: 1},
					InstructionI32Add{},
					InstructionEnd{},
				},
			},
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("unexpected module: %#v", got)
	}
}
