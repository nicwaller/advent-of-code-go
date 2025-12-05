package sugar

import (
	"iter"

	. "advent-of-code/lib/sequent/v1"
)

// This package is intended to be imported into the default namespace
// using the . operator with the import.

type S[T any] struct{}

func (S[T]) Seq(seq iter.Seq[T]) Iterator[T] {
	return SeqIterator[T](seq)
}

func (S[T]) List(list []T) Iterator[T] {
	return ListIterator(list)
}

func (S[T]) Args(args ...T) Iterator[T] {
	return ListIterator(args)
}
