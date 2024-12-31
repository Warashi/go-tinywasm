package binary

type RefType byte

const (
	RefTypeFunc   RefType = 0x70
	RefTypeExtern RefType = 0x6f
)

type TableType struct {
	ElementType RefType
	Limits      Limits
}
