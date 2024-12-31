package runtime

import (
	"fmt"
	"iter"

	"github.com/Warashi/go-tinywasm/binary"
	tbinary "github.com/Warashi/go-tinywasm/types/binary"
	"github.com/Warashi/go-tinywasm/types/runtime"
)

const PageSize = 65536 // 64 Ki

type Store struct {
	funcs    []runtime.FuncInst
	module   runtime.ModuleInst
	memories []runtime.MemoryInst
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

		locals := make([]tbinary.ValueType, 0, len(body.Locals))

		for _, local := range body.Locals {
			locals = append(locals, local.ValueType)
		}

		funcInst := runtime.InternalFuncInst{
			FuncType: funcType,
			Code: runtime.Func{
				Locals: locals,
				Body:   body.Code,
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

	for _, data := range module.DataSection() {
		memory := memories[data.MemoryIndex]
		if int(data.Offset)+len(data.Init) > len(memory.Data) {
			return nil, fmt.Errorf("data segment does not fit in memory")
		}
		copy(memory.Data[data.Offset:], data.Init)
	}

	return &Store{
		funcs:    funcs,
		memories: memories,
		module: runtime.ModuleInst{
			Exports: exports,
		},
	}, nil
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
