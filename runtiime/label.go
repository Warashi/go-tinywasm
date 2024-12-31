package runtime

type LabelKind int

const (
	LabelKindBlock LabelKind = iota
	LabelKindLoop
	LabelKindIf
)

type Label struct {
	kind           LabelKind
	programCounter int
	stackPointer   int
	arity          int
}
