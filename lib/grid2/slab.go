package grid2

import (
	"image"
	"iter"
)

// Slab is part of a Grid
type Slab[T comparable] struct {
	underlying *Grid[T]
	bounds     image.Rectangle
}

// Bounds() is provided as a way to access the column/row counts
// it returns by value, to avoid accidentally modifying the underlying boundary
func (s *Slab[T]) Bounds() image.Rectangle {
	return s.bounds
}

func (s *Slab[T]) Points() iter.Seq[image.Point] {
	return func(yield func(point image.Point) bool) {
		origin := s.bounds.Canon().Min
		terminus := s.bounds.Canon().Max

		var p image.Point
		for p.X = origin.X; p.Y < terminus.Y; p.Y++ {
			for ; p.X < terminus.X; p.X++ {
				if !yield(p) {
					return
				}
			}
		}
	}
}

func (s *Slab[T]) Cells() iter.Seq[Cell[T]] {
	return func(yield func(Cell[T]) bool) {
		for p := range s.Points() {
			c := Cell[T]{
				underlying:  s.underlying,
				coordinates: p,
			}
			if !yield(c) {
				return
			}
		}
	}
}
