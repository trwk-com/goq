package goq

type (
	Type interface {
		AcceptType(v TypeVisitor) error
	}

	TypeVisitor interface {
		VisitBoolType(t BoolType) error
		VisitDateType(t DateType) error
		VisitFloatType(t FloatType) error
		VisitIntType(t IntType) error
		VisitStringType(t StringType) error
	}
)
