package grid

type Grid[T comparable] struct {
	storage []T
	// Grid invariant: len(offsets) == len(dimensions)
	dimensions []int // ordered major to minor
	offsets    []int // ordered major to minor
	jumps      []int // pre-calculated cache of how much to jump for each dimension
}

type Cell []int

func IsZeroCell(c Cell) bool {
	for _, v := range c {
		if v != 0 {
			return false
		}
	}
	return true
}

func (grid *Grid[T]) recalculateJumps() {
	grid.jumps = make([]int, len(grid.dimensions))
	jumpSize := 1
	for d := len(grid.dimensions) - 1; d >= 0; d-- {
		grid.jumps[d] = jumpSize
		jumpSize *= grid.dimensions[d]
	}
}

// OffsetFromCell returns the index of the storage slice that corresponds to the requested cell
// 2D example, 3x3 grid:
//   0 1 2
//   3 4 5
//   6 7 8
func (grid *Grid[T]) OffsetFromCell(cell Cell) int {
	pos := 0
	for d, z := range cell {
		pos += (z - grid.offsets[d]) * grid.jumps[d]
	}
	return pos
}

func (grid *Grid[T]) CellFromOffset(offset int) Cell {
	c := make([]int, len(grid.dimensions))
	for d, j := range grid.jumps {
		c[d] = offset / j
		offset -= c[d] * j
	}
	for d, ofs := range grid.offsets {
		c[d] += ofs
	}
	return c
}

func (grid *Grid[T]) Values() *[]T {
	return &grid.storage
}
