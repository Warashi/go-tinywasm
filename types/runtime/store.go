package runtime

import "github.com/Warashi/go-tinywasm/types/binary"


type FuncInst interface {
	isFuncInst()
}

type InternalFuncInst struct {
	FuncType binary.FuncType
	Code     Func
}

func (f InternalFuncInst) isFuncInst() {}

type Func struct {
	Locals []binary.ValueType
	Body   []Instruction
}

type ExternalFuncInst struct {
	Module   string
	Func     string
	FuncType binary.FuncType
}

func (f ExternalFuncInst) isFuncInst() {}

type ExportInst struct {
	Name string
	Desc binary.ExportDesc
}

type ModuleInst struct {
	Exports map[string]ExportInst
}

func (m ModuleInst) Exported(name string) (ExportInst, bool) {
	e, ok := m.Exports[name]
	return e, ok
}

type MemoryInst struct {
	Data []byte
	Max  uint32
}
