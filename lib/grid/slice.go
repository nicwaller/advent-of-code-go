package grid

import (
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
)

type Slice []Range

type Range struct {
	origin   int // inclusive
	terminus int // exclusive
}

// Dimensions needs to be TODO: tested
func (slice Slice) Dimensions() []int {
	dim := make([]int, len(slice))
	for i, z := range slice {
		dim[i] = IntAbs(z.terminus-z.origin) + 1
	}
	return dim
}

func IntAbs(x int) int {
	if x < 0 {
		return 0 - x
	} else {
		return x
	}
}

func Line(a Cell, b Cell) iter.Iterator[Cell] {
	unitize := func(f int, g int) int {
		// can this be done more cleverly with bit shift?
		switch {
		case f > g:
			return 1
		case f == g:
			return 0
		case f < g:
			return -1
		}
		panic("impossible")
	}
	curPos := make([]int, len(a))
	nextPos := make([]int, len(a))
	copy(curPos, a)
	copy(nextPos, a)
	done := false
	return iter.Iterator[Cell]{
		Next: func() bool {
			copy(curPos, nextPos)
			for d := 0; d < len(a); d++ {
				nextPos[d] += unitize(b[d], a[d])
			}
			if done {
				return false
			}
			done = true
			for d := 0; d < len(a); d++ {
				if curPos[d] != b[d] {
					done = false
				}
			}
			return true
		},
		Value: func() Cell {
			// Yes, this copying and allocation is necessary
			// so that List() works correctly on the iterator
			ret := make([]int, len(curPos))
			copy(ret, curPos)
			return ret
		},
	}
}

func SliceEnclosing(cells ...Cell) Slice {
	nDimensions := len(cells[0])
	slice := make([]Range, nDimensions)
	// initialize
	c := cells[0]
	for d := 0; d < nDimensions; d++ {
		slice[d].origin = c[d]
		slice[d].terminus = c[d] + 1
	}
	// loop
	for _, c := range cells {
		for d := 0; d < nDimensions; d++ {
			slice[d].origin = util.IntMin(slice[d].origin, c[d])
			slice[d].terminus = util.IntMax(slice[d].terminus, c[d]+1)
		}
	}
	return slice
}

//	rNormal := r.Normalize()
//	x := rNormal.Bounds[0].X - 1
//	y := rNormal.Bounds[0].Y
//	return iter.Iterator[Coord2D]{
//		Next: func() bool {
//			x++
//			if x > rNormal.Bounds[1].X {
//				x = rNormal.Bounds[0].X
//				y++
//			}
//			if y > rNormal.Bounds[1].Y {
//				return false
//			}
//			fmt.Printf("x: %d, y: %d\n", x, y)
//			return true
//		},
//		Value: func() Coord2D {
//			return Coord2D{
//				X: x,
//				Y: y,
//			}
//		},
//	}

func (slice Slice) Origin() Cell {
	c := make([]int, len(slice))
	for i := 0; i < len(c); i++ {
		c[i] = slice[i].origin
	}
	return c
}

// Cells pretty hard to implement as an n-dimensional iterator!
func (slice Slice) Cells() iter.Iterator[Cell] {
	nDimensions := len(slice)
	_ = nDimensions
	current := slice.Origin()
	current[len(current)-1]--
	carry := func() {
		for d := len(slice) - 1; d >= 0; d-- {
			if current[d] == slice[d].terminus {
				if d == 0 {
					current = nil
					return
				}
				current[d] = slice[d].origin
				current[d-1]++
			}
		}
	}
	return iter.Iterator[Cell]{
		Next: func() bool {
			current[len(slice)-1]++
			carry()
			return current != nil
			//if current[1] == slice[1].terminus {
			//	current[1] = slice[1].origin
			//}
			//for d := 0; d < nDimensions; d++ {
			//	if current[d] > slice[0].terminus {
			//		// TODO: figure out the terminal condition
			//		current[d] = slice[0].origin
			//	}
			//}
		},
		Value: func() Cell {
			return current
		},
	}
}

func (grid *Grid[T]) Select(cells ...Cell) Slice {
	return SliceEnclosing(cells...)
}
