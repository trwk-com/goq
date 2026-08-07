package goq

type (
	BoolType struct{}
)

func (t BoolType) AcceptType(v TypeVisitor) error {
	return v.VisitBoolType(t)
}

var (
	_ Type = BoolType{}
)
