package binary

type SectionCode byte

const (
	SectionCodeCustom SectionCode = iota
	SectionCodeType
	SectionCodeImport
	SectionCodeFunction
	_
	SectionCodeMemory
	_
	SectionCodeExport
	_
	_
	SectionCodeCode
	SectionCodeData
)

type Function struct {
	locals []FunctionLocal
	code   []Instruction
}

func (f Function) Locals() []FunctionLocal {
	return f.locals
}

func (f Function) Code() []Instruction {
	return f.code
}
