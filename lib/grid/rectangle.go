package grid

import (
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
)

// TODO: should Rect be inclusive on both sides?
type Rect2D struct {
	Bounds [2]Coord2D
}

func MakeRect(x1 int, y1 int, x2 int, y2 int) Rect2D {
	return Rect2D{Bounds: [2]Coord2D{
		{
			X: x1,
			Y: y1,
		},
		{
			X: x2,
			Y: y2,
		},
	}}
}

// Normalize ensures bounds[1] > bounds[0]
func (r Rect2D) Normalize() Rect2D {
	return Rect2D{
		Bounds: [2]Coord2D{
			{
				X: util.IntMin(r.Bounds[0].X, r.Bounds[1].X),
				Y: util.IntMin(r.Bounds[0].Y, r.Bounds[1].Y),
			},
			{
				X: util.IntMax(r.Bounds[0].X, r.Bounds[1].X),
				Y: util.IntMax(r.Bounds[0].Y, r.Bounds[1].Y),
			},
		},
	}
}

// Cells returns all the coordinates enclosed within this rectangle
func (r Rect2D) Cells() iter.Iterator[Coord2D] {
	rNormal := r.Normalize()
	x := rNormal.Bounds[0].X - 1
	y := rNormal.Bounds[0].Y
	return iter.Iterator[Coord2D]{
		Next: func() bool {
			x++
			if x > rNormal.Bounds[1].X {
				x = rNormal.Bounds[0].X
				y++
			}
			if y > rNormal.Bounds[1].Y {
				return false
			}
			//fmt.Printf("x: %d, y: %d\n", x, y)
			return true
		},
		Value: func() Coord2D {
			return Coord2D{
				X: x,
				Y: y,
			}
		},
	}
}

func RectUnion(rects ...Rect2D) Rect2D {
	union := rects[0].Normalize()
	for _, r := range rects {
		next := r.Normalize()
		union.Bounds[0].X = util.IntMin(union.Bounds[0].X, next.Bounds[0].X)
		union.Bounds[0].Y = util.IntMin(union.Bounds[0].Y, next.Bounds[0].Y)
		union.Bounds[1].X = util.IntMax(union.Bounds[1].X, next.Bounds[1].X)
		union.Bounds[1].Y = util.IntMax(union.Bounds[1].Y, next.Bounds[1].Y)
	}
	return union
}
