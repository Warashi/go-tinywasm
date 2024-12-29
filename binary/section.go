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
