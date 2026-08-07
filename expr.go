package goq

type (
	Expr interface {
		Type() Type
		AcceptExpr(v ExprVisitor) error
	}

	ExprVisitor interface {
		VisitAttrExpr(e AttrExpr) error
	}
)
