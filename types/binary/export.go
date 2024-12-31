package binary

type Export struct {
	Name string
	Desc ExportDesc
}

type ExportDesc interface {
	isExportDesc()
}

type ExportDescFunc struct {
	Index uint32
}

func (e ExportDescFunc) isExportDesc() {}

type ExportDescTable struct {
	Index uint32
}

func (e ExportDescTable) isExportDesc() {}

type ExportDescMemory struct {
	Index uint32
}

func (e ExportDescMemory) isExportDesc() {}

type ExportDescGlobal struct {
	Index uint32
}

func (e ExportDescGlobal) isExportDesc() {}
