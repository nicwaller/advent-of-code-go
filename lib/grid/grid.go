package grid

import (
	"advent-of-code/lib/iter"
	"fmt"
	"strconv"
	"strings"
)

type Grid[T comparable] struct {
	storage    []T
	dimensions []int
}

func FromString(s string) Grid[string] {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	height := len(lines)
	width := len(lines[0])
	storage := make([]string, width*height)
	offset := 0
	for _, line := range lines {
		// TODO: validate line length
		for _, val := range line {
			storage[offset] = string(val)
			offset++
		}
	}
	return Grid[string]{
		storage:    storage,
		dimensions: []int{height, width},
	}
}

func FromDelimitedString(s string, delim rune) Grid[string] {
	split := func(line string) []string {
		return strings.FieldsFunc(line, func(r rune) bool {
			return r == delim
		})
	}
	lines := strings.Split(strings.TrimSpace(s), "\n")
	height := len(lines)
	width := len(split(lines[0]))
	storage := make([]string, width*height)
	offset := 0
	for _, line := range lines {
		for _, val := range split(line) {
			storage[offset] = val
			offset++
		}
	}
	return Grid[string]{
		storage:    storage,
		dimensions: []int{height, width},
	}
}

func FromStringAsInt(s string) Grid[int] {
	unsafeAtoi := func(s string) int {
		res, _ := strconv.Atoi(s)
		return res
	}
	return TransformGrid(FromString(s), unsafeAtoi)
}

func FromDelimitedStringAsInt(s string, delim rune) Grid[int] {
	unsafeAtoi := func(s string) int {
		res, _ := strconv.Atoi(s)
		return res
	}
	return TransformGrid(FromDelimitedString(s, delim), unsafeAtoi)
}

// TODO: IntGridFromString

func NewGrid[T comparable](dimensions ...int) Grid[T] {
	size := 1
	for _, d := range dimensions {
		size *= d
	}
	return Grid[T]{
		storage:    make([]T, size),
		dimensions: dimensions,
	}
}

func (grid *Grid[T]) String() string {
	switch len(grid.dimensions) {
	case 1:
		return fmt.Sprint(grid.storage)
	case 2:
		rowLength := grid.dimensions[1]
		rows := len(grid.storage) / rowLength
		var rowStrings []string
		for row := 0; row < rows; row++ {
			offset := row * rowLength
			line := fmt.Sprint(grid.storage[offset : offset+rowLength])
			rowStrings = append(rowStrings, line)
		}
		return strings.Join(rowStrings, "\n")
	default:
		return fmt.Sprintf("Cannot print %d-dimensional grid", len(grid.dimensions))
	}
	// be mindful of dimensions
	// maybe redo this with reader/writer
}

func (grid *Grid[T]) Fill(fillValue T) {
	for i, _ := range grid.storage {
		grid.storage[i] = fillValue
	}
}

func (grid *Grid[T]) Values() *[]T {
	return &grid.storage
}

func (grid *Grid[T]) Map(fn func(v T) T) {
	for i, item := range grid.storage {
		grid.storage[i] = fn(item)
	}
}

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

func (grid *Grid[T]) ColumnIter() iter.Iterator[*[]T] {
	// Assumes RowMajor storage
	if len(grid.dimensions) != 2 {
		panic("RowIterator() only makes sense for 2D grids")
	}
	colIndex := 0
	var column []T
	return iter.Iterator[*[]T]{
		Next: func() bool {
			if colIndex < grid.ColumnCount() {
				column = grid.Column(colIndex)
				colIndex++
				return true
			} else {
				return false
			}
		},
		Value: func() *[]T {
			return &column
		},
	}
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

func (grid *Grid[T]) Transpose() Grid[T] {
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

// Change type of grid (map each element)

//FindFunc() // return set of coords
//Bounds() Rect2D
//Get(x,y,z) / Set(x,y,z)
//Frequency() // most common elements
//Add() // within given bounds, inclusive
//Zero() or Fill () with custom value
//Neighbours() / NeighboursManhattan // horizontal, vertical, diagonal, of any size, also respecting bounds
//ManhattanDistance()
//String() // SPrettyPribt
//GridIterator()
//FloodFill
//A*

// TODO: origin can be 0,0 or 1,1
