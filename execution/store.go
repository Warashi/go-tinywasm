package execution

import (
	"iter"

	"github.com/Warashi/go-tinywasm/binary"
)

type Store struct {
	funcs []FuncInst
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

func NewStore(module *binary.Module) (*Store, error) {
	var funcs []FuncInst

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

	return &Store{
		funcs: funcs,
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
