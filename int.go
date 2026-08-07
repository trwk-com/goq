package goq

type (
	IntType struct {
		min int64
		max int64
	}
)

func (t IntType) AcceptType(v TypeVisitor) error {
	return v.VisitIntType(t)
}

var (
	_ Type = IntType{}
)
