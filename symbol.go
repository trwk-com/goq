package goq

type (
	Symbol interface {
		Name() string
		AcceptSymbol(v SymbolVisitor) error
	}

	SymbolVisitor interface {
		VisitAttrSymbol(s Attr) error
		VisitEdgeSymbol(s Edge) error
		VisitEntitySymbol(s Entity) error
	}
)
