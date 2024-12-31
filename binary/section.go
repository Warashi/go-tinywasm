package binary

type SectionCode byte

const (
	SectionCodeCustom SectionCode = iota
	SectionCodeType
	SectionCodeImport
	SectionCodeFunction
	SectionCodeTable
	SectionCodeMemory
	SectionCodeGlobal
	SectionCodeExport
	SectionCodeStart
	SectionCodeElement
	SectionCodeCode
	SectionCodeData
	SectionCodeDataCount
)
