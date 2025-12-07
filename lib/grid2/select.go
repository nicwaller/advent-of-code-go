package grid2

import "iter"

func (g *Grid[T]) Select() Selector[T] {
	return Selector[T]{
		underlying: g,
	}
}

type Selector[T comparable] struct {
	underlying *Grid[T]
}

func (s Selector[T]) ByValue(predicate func(T) bool) iter.Seq[Cell[T]] {
	g := s.underlying
	return func(yield func(Cell[T]) bool) {
		for i, v := range g.storage {
			if predicate(v) {
				c := Cell[T]{
					underlying:  g,
					coordinates: getPoint(g, i),
				}
				if !yield(c) {
					return
				}
			}
		}
	}
}

func (s Selector[T]) ValueEquals(target T) iter.Seq[Cell[T]] {
	return s.ByValue(func(v T) bool {
		return v == target
	})
}
