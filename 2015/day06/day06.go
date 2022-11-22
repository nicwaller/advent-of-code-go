package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2015, 6)
	aoc.TestLiteral(run, "turn on 0,0 through 999,999", "1000000", "")
	aoc.TestLiteral(run, "toggle 0,0 through 999,0", "1000", "")
	aoc.TestLiteral(run, "toggle 0,0 through 999,999", "", "2000000")
	aoc.Test(run, "input.txt", "377891", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	g := grid.NewGrid[bool](1000, 1000)
	g2 := grid.NewGrid[int](1000, 1000)
	for _, line := range aoc.InputLines() {
		f := util.NumberFields(line)
		s := grid.SliceEnclosing(
			grid.Cell{f[0], f[1]},
			grid.Cell{f[2], f[3]},
		)
		if strings.HasPrefix(line, "toggle") {
			g.MapSlice(func(b bool) bool { return !b }, s)
		} else {
			b := map[string]bool{
				"on":  true,
				"off": false,
			}[strings.Fields(line)[1]]
			g.FillSlice(b, s)
		}
		switch {
		case strings.HasPrefix(line, "turn on"):
			g2.MapSlice(func(b int) int { return b + 1 }, s)
		case strings.HasPrefix(line, "turn off"):
			g2.MapSlice(func(b int) int { return util.IntMax(0, b-1) }, s)
		case strings.HasPrefix(line, "toggle"):
			g2.MapSlice(func(b int) int { return b + 2 }, s)
		default:
			panic(line)
		}
	}
	*p1 = strconv.Itoa(g.Filter(func(c grid.Cell, b bool) bool { return b }).Count())
	*p2 = strconv.Itoa(iter.ListIterator(g2.Values()).Reduce(util.IntSum, 0))
}
