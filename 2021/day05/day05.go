package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2021, 5)
	aoc.Test(run, "sample.txt", "5", "12")
	aoc.Test(run, "input.txt", "5585", "17193")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(part1())
	*p2 = strconv.Itoa(part2())
}

type fileType []grid.Cell

func parseFile() fileType {
	clouds := make([]grid.Cell, 0)
	for lines := aoc.InputLinesIterator(); lines.Next(); {
		f := util.NumberFields(lines.Value())
		clouds = append(clouds,
			grid.Cell{f[0], f[1]},
			grid.Cell{f[2], f[3]},
		)
	}
	return clouds
}

func part1() int {
	input := parseFile()
	extents := grid.SliceEnclosing(input...)
	g := grid.NewGridFromSlice[int](extents)
	incr := func(i int) int { return i + 1 }
	for i := 0; i < len(input); i += 2 {
		// skip lines that aren't horizontal or vertical
		if !(input[i][0] == input[i+1][0] || input[i][1] == input[i+1][1]) {
			continue
		}
		g.MapIter(incr, grid.Line(input[i], input[i+1]))
	}
	return g.Filter(func(v int) bool { return v >= 2 }).Count()
}

func part2() int {
	input := parseFile()
	extents := grid.SliceEnclosing(input...)
	g := grid.NewGridFromSlice[int](extents)
	incr := func(i int) int { return i + 1 }
	for i := 0; i < len(input); i += 2 {
		g.MapIter(incr, grid.Line(input[i], input[i+1]))
	}
	return g.Filter(func(v int) bool { return v >= 2 }).Count()
}
