package grid

import (
	"strconv"
	"strings"
)

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
	g := Grid[string]{
		storage:    storage,
		dimensions: []int{height, width},
		offsets:    make([]int, 2),
	}
	g.recalculateJumps()
	return g
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
	g := Grid[string]{
		storage:    storage,
		dimensions: []int{height, width},
		offsets:    make([]int, 2),
	}
	g.recalculateJumps()
	return g
}

func FromStringAsInt(s string) Grid[int] {
	orig := FromString(s)
	unsafeAtoi := func(s string) int {
		res, _ := strconv.Atoi(s)
		return res
	}
	return TransformGrid(orig, unsafeAtoi)
}

func FromDelimitedStringAsInt(s string, delim rune) Grid[int] {
	unsafeAtoi := func(s string) int {
		res, _ := strconv.Atoi(s)
		return res
	}
	return TransformGrid(FromDelimitedString(s, delim), unsafeAtoi)
}

func (grid Grid[T]) Copy() Grid[T] {
	return Copy(grid)
}

func Copy[T comparable](old Grid[T]) Grid[T] {
	newG := Grid[T]{}

	newG.storage = make([]T, len(old.storage))
	copy(newG.storage, old.storage)

	newG.dimensions = make([]int, len(old.dimensions))
	copy(newG.dimensions, old.dimensions)

	newG.offsets = make([]int, len(old.offsets))
	copy(newG.offsets, old.offsets)

	newG.recalculateJumps()
	return newG
}

// NewGrid produces a grid of requested size with a zero-point Origin (0, 0, ...)
func NewGrid[T comparable](dimensions ...int) Grid[T] {
	size := 1
	for _, d := range dimensions {
		size *= d
	}
	g := Grid[T]{
		storage:    make([]T, size),
		dimensions: dimensions,
		offsets:    make([]int, len(dimensions)),
	}
	g.recalculateJumps()
	return g
}

func NewGridFromSlice[T comparable](slice Slice) Grid[T] {
	g := NewGrid[T](slice.Dimensions()...)
	for i, _ := range slice {
		g.offsets[i] = slice[i].Origin
	}
	return g
}
