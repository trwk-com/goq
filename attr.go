package goq

type (
	Attr interface {
		Member
		Type() Type
		Required() bool
	}

	AttrExpr struct {
		path []Edge
		attr Attr
	}
)

func (e AttrExpr) Type() Type {
	return e.attr.Type()
}

func (e AttrExpr) AcceptExpr(v ExprVisitor) error {
	return v.VisitAttrExpr(e)
}

var (
	_ Expr   = AttrExpr{}
)
