package runtime

import (
	"fmt"
	"iter"

	"github.com/Warashi/go-tinywasm/binary"
)

const PageSize = 65536 // 64 Ki

type Store struct {
	funcs    []FuncInst
	module   ModuleInst
	memories []MemoryInst
}

type FuncInst interface {
	isFuncInst()
}

type InternalFuncInst struct {
	funcType binary.FuncType
	code     Func
}

type Func struct {
	locals []binary.ValueType
	body   []binary.Instruction
}

func (f InternalFuncInst) isFuncInst() {}

type ExternalFuncInst struct {
	module   string
	fn       string
	funcType binary.FuncType
}

func (f ExternalFuncInst) isFuncInst() {}

type ExportInst struct {
	name string
	desc binary.ExportDesc
}

type ModuleInst struct {
	exports map[string]ExportInst
}

type MemoryInst struct {
	data []byte
	max  uint32
}

func NewStore(module *binary.Module) (*Store, error) {
	var funcs []FuncInst

	for _, impt := range module.ImportSection() {
		moduleName := impt.Module()
		field := impt.Field()
		switch desc := impt.Desc().(type) {
		case binary.ImportDescFunc:
			if desc.Index() < 0 || len(module.TypeSection()) <= int(desc.Index()) {
				return nil, fmt.Errorf("invalid function index: %d", desc.Index())
			}
			funcType := module.TypeSection()[desc.Index()]

			funcs = append(funcs, ExternalFuncInst{
				module:   moduleName,
				fn:       field,
				funcType: funcType,
			})
		}
	}

	for body, typeIdx := range zipSlice(module.CodeSection(), module.FunctionSection()) {
		funcType := module.TypeSection()[typeIdx]

		locals := make([]binary.ValueType, 0, len(body.Locals()))

		for _, local := range body.Locals() {
			locals = append(locals, local.ValueType())
		}

		funcInst := InternalFuncInst{
			funcType: funcType,
			code: Func{
				locals: locals,
				body:   body.Code(),
			},
		}

		funcs = append(funcs, funcInst)
	}

	exports := make(map[string]ExportInst, len(module.ExportSection()))
	for _, export := range module.ExportSection() {
		exports[export.Name()] = ExportInst{
			name: export.Name(),
			desc: export.Desc(),
		}
	}

	memories := make([]MemoryInst, 0, len(module.MemorySection()))
	for _, memory := range module.MemorySection() {
		mem := MemoryInst{
			data: make([]byte, memory.Limits().Min()*PageSize),
			max:  memory.Limits().Max(),
		}
		memories = append(memories, mem)
	}

	for _, data := range module.DataSection() {
		memory := memories[data.MemoryIndex()]
		if int(data.Offset())+len(data.Init()) > len(memory.data) {
			return nil, fmt.Errorf("data segment does not fit in memory")
		}
		copy(memory.data[data.Offset():], data.Init())
	}

	return &Store{
		funcs:    funcs,
		memories: memories,
		module: ModuleInst{
			exports: exports,
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
