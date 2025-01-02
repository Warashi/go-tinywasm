package runtime

import (
	"fmt"
	"iter"

	"github.com/Warashi/wasmium/binary"
	tbinary "github.com/Warashi/wasmium/types/binary"
	"github.com/Warashi/wasmium/types/instruction"
	"github.com/Warashi/wasmium/types/runtime"
)

const PageSize = 65536 // 64 Ki

type Store struct {
	funcs    []runtime.FuncInst
	module   runtime.ModuleInst
	memories []runtime.MemoryInst
	globals  []runtime.GlobalInst
}

func NewStore(module *binary.Module) (*Store, error) {
	var funcs []runtime.FuncInst

	for _, impt := range module.ImportSection() {
		moduleName := impt.Module
		field := impt.Field
		switch desc := impt.Desc.(type) {
		case tbinary.ImportDescFunc:
			if desc.Index < 0 || len(module.TypeSection()) <= int(desc.Index) {
				return nil, fmt.Errorf("invalid function index: %d", desc.Index)
			}
			funcType := module.TypeSection()[desc.Index]

			funcs = append(funcs, runtime.ExternalFuncInst{
				Module:   moduleName,
				Func:     field,
				FuncType: funcType,
			})
		}
	}

	for body, typeIdx := range zipSlice(module.CodeSection(), module.FunctionSection()) {
		funcType := module.TypeSection()[typeIdx]

		localslen := 0

		for _, local := range body.Locals {
			localslen += int(local.TypeCount)
		}
		locals := make([]tbinary.ValueType, 0, localslen)

		for _, local := range body.Locals {
			for range local.TypeCount {
				locals = append(locals, local.ValueType)
			}
		}

		insts, err := instruction.Convert(body.Code)
		if err != nil {
			return nil, fmt.Errorf("failed to convert instructions: %w", err)
		}

		funcInst := runtime.InternalFuncInst{
			FuncType: funcType,
			Code: runtime.Func{
				Locals: locals,
				Body:   insts,
			},
		}

		funcs = append(funcs, funcInst)
	}

	exports := make(map[string]runtime.ExportInst, len(module.ExportSection()))
	for _, export := range module.ExportSection() {
		exports[export.Name] = runtime.ExportInst{
			Name: export.Name,
			Desc: export.Desc,
		}
	}

	memories := make([]runtime.MemoryInst, 0, len(module.MemorySection()))
	for _, memory := range module.MemorySection() {
		mem := runtime.MemoryInst{
			Data: make([]byte, memory.Limits.Min*PageSize),
			Max:  memory.Limits.Max,
		}
		memories = append(memories, mem)
	}

	globals := make([]runtime.GlobalInst, 0, len(module.GlobalSection()))
	for _, global := range module.GlobalSection() {
		var v runtime.Value
		switch expr := global.InitExpr.(type) {
		case tbinary.ExprValueConstI32:
			v = runtime.ValueI32(expr)
		case tbinary.ExprValueConstI64:
			v = runtime.ValueI64(expr)
		case tbinary.ExprValueConstF32:
			v = runtime.ValueF32(expr)
		case tbinary.ExprValueConstF64:
			v = runtime.ValueF64(expr)
		default:
			return nil, fmt.Errorf("unsupported global type: %T", expr)
		}

		globals = append(globals, runtime.GlobalInst{
			Value:   v,
			Mutable: global.Type.Mutable,
		})
	}

	eval := func(expr tbinary.Expr) (int, error) {
		switch expr := expr.(type) {
		case tbinary.ExprValue:
			return expr.Int(), nil
		case tbinary.ExprGlobalIndex:
			if expr < 0 || len(globals) <= int(expr) {
				return 0, fmt.Errorf("invalid global index: %d", expr)
			}
			return globals[expr].Value.Int(), nil
		default:
			return 0, fmt.Errorf("unsupported global type: %T", expr)
		}
	}

	for _, data := range module.DataSection() {
		memory := memories[data.MemoryIndex]
		offset, err := eval(data.Offset)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate offset: %w", err)
		}
		if offset+len(data.Init) > len(memory.Data) {
			return nil, fmt.Errorf("data segment does not fit in memory")
		}
		copy(memory.Data[offset:], data.Init)
	}

	return &Store{
		funcs:    funcs,
		memories: memories,
		globals:  globals,
		module: runtime.ModuleInst{
			Exports: exports,
		},
	}, nil
}

func (s *Store) Funcs() []runtime.FuncInst {
	return s.funcs
}

func (s *Store) Module() runtime.ModuleInst {
	return s.module
}

func (s *Store) Memory(n int) (runtime.MemoryInst, error) {
	if n < 0 || len(s.memories) <= n {
		return runtime.MemoryInst{}, fmt.Errorf("invalid memory index: %d", n)
	}
	return s.memories[n], nil
}

func zipSlice[A, B any, SA ~[]A, SB ~[]B](a SA, b SB) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		for i := range min(len(a), len(b)) {
			if !yield(a[i], b[i]) {
				return
			}
		}
	}
}
