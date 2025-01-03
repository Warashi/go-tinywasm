package binary

type Expr interface {
	isExpr()
}

type ExprValue interface {
	isExprValue()
	Int() int
}

type ExprValueConstI32 int32

func (ExprValueConstI32) isExpr()      {}
func (ExprValueConstI32) isExprValue() {}
func (e ExprValueConstI32) Int() int   { return int(e) }

type ExprValueConstI64 int64

func (ExprValueConstI64) isExpr()      {}
func (ExprValueConstI64) isExprValue() {}
func (e ExprValueConstI64) Int() int   { return int(e) }

type ExprValueConstF32 [4]byte

func (ExprValueConstF32) isExpr()      {}
func (ExprValueConstF32) isExprValue() {}
func (ExprValueConstF32) Int() int     { panic("int for f32 is not allowed") }

type ExprValueConstF64 [8]byte

func (ExprValueConstF64) isExpr()      {}
func (ExprValueConstF64) isExprValue() {}
func (e ExprValueConstF64) Int() int   { panic("int for f64 is not allowed") }

type ExprGlobalIndex uint32

func (e ExprGlobalIndex) isExpr() {}
