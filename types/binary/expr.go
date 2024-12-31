package binary

type Expr interface {
	isExpr()
}

type ExprValue interface {
	isExprValue()
	Int() int
}

type ExprValueConstI32 int32

func (e ExprValueConstI32) isExpr()      {}
func (e ExprValueConstI32) isExprValue() {}
func (e ExprValueConstI32) Int() int     { return int(e) }

type ExprValueConstI64 int64

func (e ExprValueConstI64) isExpr()      {}
func (e ExprValueConstI64) isExprValue() {}
func (e ExprValueConstI64) Int() int     { return int(e) }

type ExprValueConstF32 float32

func (e ExprValueConstF32) isExpr()      {}
func (e ExprValueConstF32) isExprValue() {}
func (e ExprValueConstF32) Int() int     { return int(e) }

type ExprValueConstF64 float64

func (e ExprValueConstF64) isExpr()      {}
func (e ExprValueConstF64) isExprValue() {}
func (e ExprValueConstF64) Int() int     { return int(e) }

type ExprGlobalIndex uint32

func (e ExprGlobalIndex) isExpr() {}
