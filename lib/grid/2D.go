package grid

import "advent-of-code/lib/iter"

type Coord2D struct {
	X int
	Y int
}

func (grid *Grid[T]) RowSize() int {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("RowSize() only makes sense for 2D grids")
	}
	return grid.dimensions[1]
}

func (grid *Grid[T]) RowCount() int {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("RowCount() only makes sense for 2D grids")
	}
	return grid.dimensions[0]
}

func (grid *Grid[T]) Row(rowIndex int) []T {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("Row() only makes sense for 2D grids")
	}
	if rowIndex >= grid.RowCount() {
		panic("invalid row index")
	}
	offset := rowIndex * grid.RowSize()
	return grid.storage[offset : offset+grid.RowSize()]
}

func (grid *Grid[T]) ColumnSize() int {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("ColumnSize() only makes sense for 2D grids")
	}
	return grid.dimensions[0]
}

func (grid *Grid[T]) ColumnCount() int {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("ColumnCount() only makes sense for 2D grids")
	}
	return grid.dimensions[1]
}

func (grid *Grid[T]) Column(colIndex int) []T {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("Column() only makes sense for 2D grids")
	}
	if colIndex >= grid.ColumnCount() {
		panic("invalid column index")
	}
	column := make([]T, grid.ColumnSize())
	offset := 0
	for r := 0; r < grid.RowCount(); r++ {
		column[r] = grid.storage[offset+colIndex]
		offset += grid.RowSize()
	}
	return column
}

func (grid *Grid[T]) Width() int {
	if len(grid.dimensions) != 2 {
		panic("Width() only makes sense for 2D grids")
	}
	return grid.ColumnCount()
}

func (grid *Grid[T]) Height() int {
	if len(grid.dimensions) != 2 {
		panic("Height() only makes sense for 2D grids")
	}
	return grid.RowCount()
}

// FIXME: fix transpose (transpose+print = panic)
func (grid *Grid[T]) Transpose() Grid[T] {
	// TODO: implement this more cheaply?
	// just invert the dimensions and return a new grid
	// either leave the storage in place, or make a linear copy.
	if len(grid.dimensions) != 2 {
		panic("Transpose() only makes sense for 2D grids")
	}
	newDims := []int{grid.dimensions[1], grid.dimensions[0]}
	newStorage := make([]T, len(grid.storage))
	writeIndex := 0
	for oldX := 0; oldX < grid.Width(); oldX++ {
		for oldY := 0; oldY < grid.Height(); oldY++ {
			readIndex := (oldY * grid.RowSize()) + oldX
			newStorage[writeIndex] = grid.storage[readIndex]
			writeIndex++
		}
	}
	return Grid[T]{
		storage:    newStorage,
		dimensions: newDims,
	}
}

func (grid *Grid[T]) RowIter() iter.Iterator[[]T] {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("RowIterator() only makes sense for 2D grids")
	}
	rowIndex := 0
	var row []T
	return iter.Iterator[[]T]{
		Next: func() bool {
			if rowIndex < grid.RowCount() {
				row = grid.Row(rowIndex)
				rowIndex++
				return true
			} else {
				return false
			}
		},
		Value: func() []T {
			return row
		},
	}
}

//func (grid *Grid[T]) RowIterator() func() []T {
//	// Assumes RowMajor storage
//	if len(grid.dimensions) != 2 {
//		panic("RowIterator() only makes sense for 2D grids")
//	}
//	rowIndex := 0
//	return func() []T {
//		if rowIndex < grid.RowCount() {
//			row := grid.Row(rowIndex)
//			rowIndex++
//			return row
//		} else {
//			return nil
//		}
//	}
//}

//func (grid *Grid[T]) ColumnIterator() func() []T {
//	if len(grid.dimensions) != 2 {
//		panic("RowIterator() only makes sense for 2D grids")
//	}
//	colIndex := 0
//	return func() []T {
//		if colIndex < grid.ColumnCount() {
//			row := grid.Column(colIndex)
//			colIndex++
//			return row
//		} else {
//			return nil
//		}
//	}
//}

func (grid *Grid[T]) ColumnIter() iter.Iterator[[]T] {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("RowIterator() only makes sense for 2D grids")
	}
	colIndex := 0
	var column []T
	return iter.Iterator[[]T]{
		Next: func() bool {
			if colIndex < grid.ColumnCount() {
				column = grid.Column(colIndex)
				colIndex++
				return true
			} else {
				return false
			}
		},
		Value: func() []T {
			return column
		},
	}
}
