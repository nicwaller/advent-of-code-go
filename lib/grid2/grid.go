package grid2

import (
	"fmt"
	"image"
	"iter"
	"strings"
)

type Grid[T comparable] struct {
	// storage is always row-major
	storage []T
	bounds  image.Rectangle // expected to always be well-formed
}

func NewGrid[T comparable](width, height int) *Grid[T] {
	bounds := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: width, Y: height},
	}
	return NewGridWithBounds[T](bounds)
}

func NewGridWithBounds[T comparable](bounds image.Rectangle) *Grid[T] {
	return &Grid[T]{
		storage: make([]T, bounds.Dx()*bounds.Dy()),
		bounds:  bounds,
	}
}

// Y starts at 0 with the first line and increases downwards
// the origin (0,0) is at the top-left of the file/string
func NewGridFromString(s string) *Grid[rune] {
	// avoid having a row of nulls at the bottom of the grid
	// because unix text files normally end with a newline
	s = strings.TrimRight(s, "\n")

	lines := strings.Split(s, "\n")
	height := len(lines)
	width := -1
	for _, line := range lines {
		width = max(len(line), width)
	}
	g := NewGrid[rune](width, height)

	for y, line := range lines {
		for x, r := range line {
			g.Set(image.Point{X: x, Y: y}, r)
		}
	}

	return g
}

func getOffset[T comparable](g *Grid[T], p image.Point) int {
	if !p.In(g.bounds) {
		panic(fmt.Sprintf("cell %v out of bounds %v", p, g.bounds))
	}
	rel := p.Sub(g.bounds.Canon().Min)
	offset := g.bounds.Dx()*rel.Y + rel.X
	if offset >= len(g.storage) {
		// this should be impossible given the previous checks passed
		panic(fmt.Errorf("offset out of bounds"))
	}
	return offset
}

// I think I might want this later. -NW
//
//goland:noinspection GoUnusedFunction
func getPoint[T comparable](g *Grid[T], offset int) image.Point {
	if offset < 0 {
		panic(fmt.Errorf("negative offset"))
	}
	if offset >= len(g.storage) {
		panic(fmt.Errorf("offset out of bounds"))
	}
	rowLength := g.bounds.Dx()
	return image.Point{
		X: offset % rowLength,
		Y: offset / rowLength,
	}
}

// PERF: this is safe but inefficient
func (g *Grid[T]) Get(p image.Point) T {
	offset := getOffset(g, p)
	return g.storage[offset]
}

// PERF: this is safe but inefficient
func (g *Grid[T]) Set(p image.Point, v T) {
	offset := getOffset(g, p)
	g.storage[offset] = v
}

func (g *Grid[T]) All() *Slab[T] {
	return &Slab[T]{
		underlying: g,
		bounds:     g.bounds,
	}
}

// this is indexed starting from 0, not using the coordinate grid
func (g *Grid[T]) rowByIndex(rowIndex int) []T {
	o1 := g.bounds.Dx() * (rowIndex + 0)
	o2 := g.bounds.Dx() * (rowIndex + 1)
	return g.storage[o1:o2]
}

func (g *Grid[T]) Rows() iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for rowIndex := 0; rowIndex < g.bounds.Dy(); rowIndex++ {
			if !yield(g.rowByIndex(rowIndex)) {
				return
			}
		}
	}
}

func (g *Grid[T]) Transpose() *Grid[T] {
	tg := NewGridWithBounds[T](transposeRect(g.bounds))
	for p := range tg.All().Points() {
		dst := p
		src := image.Point{
			X: dst.Y,
			Y: dst.X,
		}
		// PERF: the Get/Set combo is fairly inefficient, but it's also simple and safe
		// transpose could be implemented faster, but I don't use it often enough to care.
		tg.Set(dst, g.Get(src))
	}
	return tg
}

func transposeRect(r image.Rectangle) image.Rectangle {
	r = r.Canon()
	return image.Rectangle{
		// preserve the same origin
		// we do not transpose around [0,0]
		Min: r.Min,
		Max: image.Point{
			X: r.Max.Y,
			Y: r.Max.Y,
		},
	}
}
