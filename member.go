package goq

type (
	Member interface {
		Symbol
		Scope() Entity
	}
)
