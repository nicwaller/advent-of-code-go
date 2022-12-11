package grid

import (
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"fmt"
)

// TODO: FillRect

func (grid *Grid[T]) Fill(fillValue T) {
	for i, _ := range grid.storage {
		grid.storage[i] = fillValue
	}
}

func (grid *Grid[T]) FillSlice(fillValue T, s Slice) {
	cells := s.Cells()
	var offset int
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Caught fatal panic in grid.FillSlice()")
			fmt.Printf("  fillValue = %v\n", fillValue)
			fmt.Printf("  slice = %v\n", s)
			fmt.Printf("  current cell = %v\n", cells.Value())
			fmt.Printf("  offset = %v\n", offset)
			panic(1)
		}
	}()
	for cells.Next() {
		v := cells.Value()
		offset = grid.OffsetFromCell(v)
		grid.storage[offset] = fillValue
	}
}

func (grid *Grid[T]) MapIter(mapFn func(v T) T, cellIter iter.Iterator[Cell]) {
	for cellIter.Next() {
		offset := grid.OffsetFromCell(cellIter.Value())
		grid.storage[offset] = mapFn(grid.storage[offset])
	}
	//for i, item := range grid.storage {
	//	grid.storage[i] = mapFn(item)
	//}
}

func (grid *Grid[T]) MapAll(fn func(v T) T) {
	for i, item := range grid.storage {
		grid.storage[i] = fn(item)
	}
}

func (grid *Grid[T]) MapSlice(mapFn func(v T) T, s Slice) {
	for cells := s.Cells(); cells.Next(); {
		offset := grid.OffsetFromCell(cells.Value())
		grid.storage[offset] = mapFn(grid.storage[offset])
	}
}

//func (grid *Grid[T]) RectMap(r Rect2D, fn func(v T) T) {
//	for cells := r.Cells(); cells.Next(); {
//		// still assuming (0,0) Origin and row-major storage
//		cell := cells.Value()
//		fmt.Println(cell)
//		offset := cell.Y*grid.RowSize() + cell.X
//		grid.storage[offset] = fn(grid.storage[offset])
//	}
//	for i, item := range grid.storage {
//		grid.storage[i] = fn(item)
//	}
//}

func TransformGrid[T comparable, Z comparable](g Grid[T], transform func(val T) Z) Grid[Z] {
	newStorage := make([]Z, len(g.storage))
	for index, item := range g.storage {
		newStorage[index] = transform(item)
	}
	return Grid[Z]{
		storage:    newStorage,
		dimensions: g.dimensions,
		offsets:    g.offsets,
		jumps:      g.jumps,
	}
}

func (grid *Grid[T]) FillFunc2D(fn func(v T, x int, y int) T) {
	rowSize := grid.RowSize()
	for i, _ := range grid.storage {
		x := i % rowSize
		y := i / rowSize
		grid.storage[i] = fn(grid.storage[i], x, y)
	}
}

// TODO: also implement FillFunc() for n-dimensions (variadic or array?)

func (grid *Grid[T]) Replace(needle T, replacement T) int {
	replacements := 0
	for index, item := range grid.storage {
		if item == needle {
			grid.storage[index] = replacement
			replacements++
		}
	}
	return replacements
}

// FloodFill returns number of changed cells
func (grid *Grid[T]) FloodFill(origin Cell, selFn func(v T) bool, fill T) int {
	sel := grid.FloodSelect(origin, selFn)
	for _, cell := range sel {
		grid.storage[grid.OffsetFromCell(cell)] = fill
	}
	return len(sel)
}

func (grid *Grid[T]) Grow(by int, emptyFill T) {
	newDims := make([]int, len(grid.dimensions))
	copy(newDims, grid.dimensions)
	for d, _ := range newDims {
		newDims[d] += by * 2
	}

	newOffsets := make([]int, len(grid.offsets))
	copy(newOffsets, grid.offsets)
	for d, _ := range newOffsets {
		newOffsets[d] -= by
	}

	sLen := f8l.Reduce(newDims, 1, util.IntProduct)
	newStorage := make([]T, sLen)

	//copy(newStorage, grid.storage)
	tmpG := Grid[T]{
		dimensions: newDims,
		offsets:    newOffsets,
		storage:    newStorage,
	}
	tmpG.Fill(emptyFill)
	tmpG.recalculateJumps()
	for cellIter := grid.Cells(); cellIter.Next(); {
		cell := cellIter.Value()
		tmpG.Set(cell, grid.Get(cell))
	}

	grid.dimensions = newDims
	grid.offsets = newOffsets
	grid.storage = newStorage
	grid.recalculateJumps()
}
