package goq

type (
	FloatType struct {
		min     float64
		exclMin bool
		max     float64
		exclMax bool
	}
)

func (t FloatType) AcceptType(v TypeVisitor) error {
	return v.VisitFloatType(t)
}

var (
	_ Type = FloatType{}
)
