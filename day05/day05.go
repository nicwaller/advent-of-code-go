package main

import (
	"advent-of-code/lib/assert"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

type fileType []grid.Cell

func parseFile() fileType {
	clouds := make([]grid.Cell, 0)
	for lines := util.ReadLines("input.txt"); lines.Next(); {
		f := util.NumberFields(lines.Value())
		clouds = append(clouds,
			grid.Cell{f[0], f[1]},
			grid.Cell{f[2], f[3]},
		)
	}
	return clouds
}

func part1(input fileType) int {
	extents := grid.SliceEnclosing(input...)
	g := grid.NewGridFromSlice[uint8](extents)
	incr := func(i uint8) uint8 { return i + 1 }
	for i := 0; i < len(input); i += 2 {
		// skip lines that aren't horizontal or vertical
		if !(input[i][0] == input[i+1][0] || input[i][1] == input[i+1][1]) {
			continue
		}
		g.MapIter(incr, grid.Line(input[i], input[i+1]))
	}
	overlaps := g.Filter(func(v uint8) bool { return v >= 2 }).Count()
	assert.Equal(overlaps, 5585)
	return overlaps
}

func part2(input fileType) int {
	extents := grid.SliceEnclosing(input...)
	g := grid.NewGridFromSlice[uint8](extents)
	incr := func(i uint8) uint8 { return i + 1 }
	for i := 0; i < len(input); i += 2 {
		g.MapIter(incr, grid.Line(input[i], input[i+1]))
	}
	overlaps := g.Filter(func(v uint8) bool { return v >= 2 }).Count()
	assert.Equal(overlaps, 17193)
	return overlaps
}
