package runtime

import (
	"github.com/Warashi/wasmium/types/binary"
)

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

func (m *MemoryInst) WriteAt(p []byte, off int64) (n int, err error) {
	if int64(len(m.Data)) < off+int64(len(p)) {
		return 0, ErrMemoryOutOfBounds
	}
	return copy(m.Data[off:], p), nil
}

func (m *MemoryInst) ReadAt(p []byte, off int64) (n int, err error) {
	if int64(len(m.Data)) < off+int64(len(p)) {
		return 0, ErrMemoryOutOfBounds
	}
	return copy(p, m.Data[off:]), nil
}
