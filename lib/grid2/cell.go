package grid2

import (
	"fmt"
	"image"
	"iter"

	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
)

type Cell[T comparable] struct {
	coordinates image.Point
	underlying  *Grid[T]
	// although it's tempting to cache a value here, premature optimization is the root of all evil
}

var _ fmt.Stringer = &Cell[any]{}
var _ fmt.GoStringer = &Cell[any]{}

func (c Cell[T]) Grid() *Grid[T] {
	return c.underlying

}

func (c Cell[T]) GoString() string {
	return fmt.Sprintf("[%d, %d]", c.coordinates.X, c.coordinates.Y)
}

func (c Cell[T]) String() string {
	return fmt.Sprintf("[%d, %d]", c.coordinates.X, c.coordinates.Y)
}

func (c Cell[T]) DebugString() string {
	return c.String()
}

func (c Cell[T]) Coordinates() image.Point {
	return c.coordinates
}

// PERF: this is safe but inefficient
func (c Cell[T]) Get() T {
	return c.underlying.Get(c.coordinates)
}

// PERF: this is safe but inefficient
func (c Cell[T]) Set(v T) {
	c.underlying.Set(c.coordinates, v)
}

func cellsInGrid[T comparable](g *Grid[T], points []image.Point) []Cell[T] {
	pointsInGrid := f8l.Filter(points, func(point image.Point) bool {
		return point.In(g.bounds)
	})

	return f8l.Map(pointsInGrid, func(p image.Point) Cell[T] {
		return Cell[T]{
			coordinates: p,
			underlying:  g,
		}
	})
}

// NeighboursCardinalAdjacent are the cells above, below, left, and right.
// In other words, they are the cells with a manhattan distance of 1.
// Only cells within the underlying grid are returned, so corners only have two neighbours.
// The result is a slice, not an iterator, because the size is always small.
func (c Cell[T]) NeighboursCardinalAdjacent() []Cell[T] {
	return cellsInGrid(c.underlying, []image.Point{
		c.coordinates.Add(Up),
		c.coordinates.Add(Down),
		c.coordinates.Add(Left),
		c.coordinates.Add(Right),
	})
}

// NeighboursEightWayAdjacent include all above, below, and diagonal that are in the grid
func (c Cell[T]) NeighboursEightWayAdjacent() []Cell[T] {
	return cellsInGrid(c.underlying, []image.Point{
		{X: c.coordinates.X - 1, Y: c.coordinates.Y - 1},
		{X: c.coordinates.X - 1, Y: c.coordinates.Y + 0},
		{X: c.coordinates.X - 1, Y: c.coordinates.Y + 1},
		{X: c.coordinates.X + 0, Y: c.coordinates.Y - 1},
		{X: c.coordinates.X + 0, Y: c.coordinates.Y + 0},
		{X: c.coordinates.X + 0, Y: c.coordinates.Y + 1},
		{X: c.coordinates.X + 1, Y: c.coordinates.Y - 1},
		{X: c.coordinates.X + 1, Y: c.coordinates.Y + 0},
		{X: c.coordinates.X + 1, Y: c.coordinates.Y + 1},
	})
}

func (c Cell[T]) ManhattanDistance(to Cell[T]) int {
	return ManhattanDistance(c.coordinates, to.coordinates)
}

func ManhattanDistance(a, b image.Point) int {
	d := b.Sub(a)
	return util.IntAbs(d.X) + util.IntAbs(d.Y)
}

// be wary of up/down inversion when building grids from strings
//
//goland:noinspection GoUnusedGlobalVariable
var (
	Up    = image.Point{X: +0, Y: -1}
	Down  = image.Point{X: +0, Y: +1}
	Left  = image.Point{X: -1, Y: +0}
	Right = image.Point{X: +1, Y: +0}
)

func (c Cell[T]) Neighbour(delta image.Point) Cell[T] {
	return Cell[T]{
		coordinates: c.coordinates.Add(delta),
		underlying:  c.underlying,
	}
}

// should it include origin? ...
func (c Cell[T]) Ray(delta image.Point) iter.Seq[Cell[T]] {
	tip := c.coordinates
	return func(yield func(Cell[T]) bool) {
		for {
			tip = tip.Add(delta)
			if !tip.In(c.underlying.bounds) {
				return
			}
			tipCell := Cell[T]{
				coordinates: tip,
				underlying:  c.underlying,
			}
			if !yield(tipCell) {
				return
			}
		}
	}
}
