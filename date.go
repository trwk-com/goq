package goq

type (
	DateType struct{}
)

func (t DateType) AcceptType(v TypeVisitor) error {
	return v.VisitDateType(t)
}

var (
	_ Type = DateType{}
)
