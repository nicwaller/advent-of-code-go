package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

var rows int

func main() {
	aoc.Select(2016, 18)
	rows = 3
	aoc.Test(run, "sample.txt", "6", "")
	rows = 10
	aoc.Test(run, "sample2.txt", "38", "")
	rows = 40
	//aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
	rows = 400000
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	trapSet := set.New[string]("^^.", ".^^", "^..", "..^")
	firstRow := aoc.InputString()
	w := len(firstRow)
	g := grid.NewGrid[string](rows, w+2)
	g.Fill(".") // "." is safe
	v := g.Values()
	copy(v[1:], iter.StringIterator(firstRow, 1).List())
	for y := 1; y < rows; y++ {
		for x := 1; x <= w; x++ {
			left := g.OffsetFromCell(grid.Cell{y - 1, x - 1})
			right := g.OffsetFromCell(grid.Cell{y - 1, x + 1})
			src := v[left : right+1]
			isTrap := trapSet.Contains(strings.Join(src, ""))
			t := map[bool]string{false: ".", true: "^"}[isTrap]
			g.Set(grid.Cell{y, x}, t)
		}
	}
	g.FillSlice(" ", grid.SliceEnclosing(grid.Cell{0, 0}, grid.Cell{rows - 1, 0}))
	g.FillSlice(" ", grid.SliceEnclosing(grid.Cell{0, w + 1}, grid.Cell{rows - 1, w + 1}))

	*p1 = strconv.Itoa(g.FilterByValue(util.Eq_(".")).Count())
	*p2 = strconv.Itoa(0)
}
