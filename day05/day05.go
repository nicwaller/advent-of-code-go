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

type fileType []grid.Rect2D

func parseFile() fileType {
	clouds := make([]grid.Rect2D, 0)
	for lines := util.ReadLines("input.txt"); lines.Next(); {
		f := util.NumberFields(lines.Value())
		clouds = append(clouds, grid.Rect2D{
			Bounds: [2]grid.Coord2D{
				grid.Coord2D{
					X: f[0],
					Y: f[1],
				},
				grid.Coord2D{
					X: f[2],
					Y: f[3],
				},
			},
		})
	}
	return clouds
}

func part1(input fileType) int {
	origin := grid.MakeRect(0, 0, 0, 0)
	union := grid.RectUnion(input...)
	extents := grid.RectUnion(origin, union)
	fmt.Println(extents)
	g := grid.NewGrid[int](extents.Bounds[1].X+1, extents.Bounds[1].Y+1)
	for _, cloud := range input {
		// TODO: skip diagonal lines
		x0 := cloud.Bounds[0].X
		y0 := cloud.Bounds[0].Y
		x1 := cloud.Bounds[1].X
		y1 := cloud.Bounds[1].Y
		if !(x0 == x1 || y0 == y1) {
			continue
		}
		s := grid.SliceEnclosing(grid.Cell{y0, x0}, grid.Cell{y1, x1})
		incr := func(i int) int { return i + 1 }
		g.MapSlice(incr, s)
	}
	overlaps := g.Filter(func(v int) bool { return v >= 2 }).List()
	fmt.Println(overlaps)
	fmt.Println(len(overlaps))
	assert.EqualNamed(len(overlaps), 5585, "number of overlaps")
	return len(overlaps)
}

func part2(input fileType) int {
	return -1
}
