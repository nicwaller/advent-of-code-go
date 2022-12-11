package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2016, 8)
	aoc.Test(run, "sample.txt", "", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	g := grid.NewGrid[string](6, 50)
	for _, line := range aoc.InputLines() {
		nf := util.NumberFields(line)
		if strings.HasPrefix(line, "rect") {
			slc := grid.SliceEnclosing(grid.Cell{0, 0}, grid.Cell{nf[1] - 1, nf[0] - 1})
			g.FillSlice("#", slc)
		} else if strings.HasPrefix(line, "rotate row") {
			rowY := nf[0]
			for x, v := range util.Rotate(nf[1], g.Row(rowY)) {
				g.Set(grid.Cell{rowY, x}, v)
			}
		} else if strings.HasPrefix(line, "rotate column") {
			colX := nf[0]
			for y, v := range util.Rotate(nf[1], g.Column(colX)) {
				g.Set(grid.Cell{y, colX}, v)
			}
		}
	}
	*p1 = strconv.Itoa(g.Filter(func(c grid.Cell, s string) bool { return s == "#" }).Count())
	*p2 = aoc.Debannerize(g, "#")
}
