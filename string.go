package goq

import "regexp"

type (
	StringType struct {
		minl uint64
		maxl uint64
		pat *regexp.Regexp
	}
)

func (t StringType) AcceptType(v TypeVisitor) error {
	return v.VisitStringType(t)
}

var (
	_ Type = StringType{}
)
