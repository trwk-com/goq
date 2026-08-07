package goq

type (
	Pred interface {
		Or(rhs Pred) Pred
		And(rhs Pred) Pred
		Not() Pred
		AcceptPred(v PredVisitor) error
	}

	PredVisitor interface {
		VisitOrPred(v OrPred) error
		VisitAndPred(v AndPred) error
		VisitNotPred(v NotPred) error
		VisitCmpPred(v CmpPred) error
	}

	OrPred struct {
		ops []Pred
	}

	AndPred struct {
		ops []Pred
	}

	NotPred struct {
		op Pred
	}

	CmpPred struct {
		op  CmpOp
		lop Expr
		rop Expr
	}

	CmpOp string
)

const (
	Eq  CmpOp = "="
	Neq CmpOp = "!="
	Lt  CmpOp = "<"
	Leq CmpOp = "<="
	Gt  CmpOp = ">"
	Geq CmpOp = ">="
)

func (p OrPred) Or(rhs Pred) Pred {
	return orPred(p, rhs)
}

func (p OrPred) And(rhs Pred) Pred {
	return andPred(p, rhs)
}

func (p OrPred) Not() Pred {
	return notPred(p)
}

func (p OrPred) AcceptPred(v PredVisitor) error {
	return v.VisitOrPred(p)
}

func (p AndPred) AcceptPred(v PredVisitor) error {
	return v.VisitAndPred(p)
}

func (p NotPred) AcceptPred(v PredVisitor) error {
	return v.VisitNotPred(p)
}

func (p CmpPred) AcceptPred(v PredVisitor) error {
	return v.VisitCmpPred(p)
}

func orPred(lhs Pred, rhs Pred) Pred {
	if rhs == nil {
		return lhs
	}

	lt, lok := lhs.(OrPred)
	rt, rok := rhs.(OrPred)
	ops := make([]Pred, 0, 10)

	if lok && rok {
		ops = append(lt.ops, rt.ops...)
	} else if lok {
		ops = append(lt.ops, rhs)
	} else if rok {
		ops = append([]Pred{lhs}, rt.ops...)
	} else {
		ops = append(ops, lhs, rhs)
	}

	return OrPred{ops: ops}
}

func andPred(lhs Pred, rhs Pred) Pred {

}

func notPred(op Pred) Pred {
	if top, ok := op.(NotPred); ok {
		return top.op
	}

	return NotPred{op: op}
}

var (
	_ Pred = OrPred{}
	_ Pred = AndPred{}
	_ Pred = NotPred{}
	_ Pred = CmpPred{}
)
