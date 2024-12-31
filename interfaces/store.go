package interfaces

import "github.com/Warashi/go-tinywasm/binary"

type FuncInst interface {
	isFuncInst()
}

type InternalFuncInst struct {
	funcType binary.FuncType
	code     Func
}

func (f InternalFuncInst) isFuncInst() {}

func (f InternalFuncInst) FuncType() binary.FuncType {
	return f.funcType
}

func (f InternalFuncInst) Code() Func {
	return f.code
}

type Func struct {
	locals []binary.ValueType
	body   []Instruction
}

func (f Func) Locals() []binary.ValueType {
	return f.locals
}

func (f Func) Body() []Instruction {
	return f.body
}

type ExternalFuncInst struct {
	module   string
	fn       string
	funcType binary.FuncType
}

func (f ExternalFuncInst) isFuncInst() {}

func (f ExternalFuncInst) Module() string {
	return f.module
}

func (f ExternalFuncInst) Fn() string {
	return f.fn
}

func (f ExternalFuncInst) FuncType() binary.FuncType {
	return f.funcType
}

type ExportInst struct {
	name string
	desc binary.ExportDesc
}

func (e ExportInst) Name() string {
	return e.name
}

func (e ExportInst) Desc() binary.ExportDesc {
	return e.desc
}

type ModuleInst struct {
	exports map[string]ExportInst
}

func (m ModuleInst) Exported(name string) (ExportInst, bool) {
	e, ok := m.exports[name]
	return e, ok
}

type MemoryInst struct {
	data []byte
	max  uint32
}
