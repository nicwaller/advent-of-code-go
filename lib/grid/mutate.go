package grid

// TODO: FillRect

func (grid *Grid[T]) Fill(fillValue T) {
	for i, _ := range grid.storage {
		grid.storage[i] = fillValue
	}
}

func (grid *Grid[T]) FillSlice(fillValue T, s Slice) {
	for cells := s.Cells(); cells.Next(); {
		grid.storage[grid.OffsetFromCell(cells.Value())] = fillValue
	}
}

func (grid *Grid[T]) Map(fn func(v T) T) {
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
//		// still assuming (0,0) origin and row-major storage
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
