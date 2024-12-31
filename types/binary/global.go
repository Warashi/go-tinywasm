package binary

type Global struct {
	Type     GlobalType
	InitExpr ExprValue
}

type GlobalType struct {
	ValueType ValueType
	Mutable   bool
}
