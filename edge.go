package goq

type (
	Edge interface {
		Member
		Target() Entity
	}
)


