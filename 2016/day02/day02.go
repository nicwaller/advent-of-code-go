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
	aoc.Select(2016, 2)
	aoc.Test(run, "sample.txt", "1985", "5DB3")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	g := grid.NewGrid[string](1, 1)
	g.Grow(1, "")
	g.FillFunc2D(func(v string, x int, y int) string {
		return strconv.Itoa(3*y + x + 1)
	})

	g2 := grid.FromString(`
*******
*  1  *
* 234 *
*56789*
* ABC *
*  D  *
*******`)
	g2.Print()
	pos := grid.Cell{0, 0}
	pos2 := grid.Cell{3, 1}
	motions := map[string][]int{
		"L": {0, -1},
		"R": {0, 1},
		"U": {-1, 0},
		"D": {1, 0},
	}
	var sb strings.Builder
	var sb2 strings.Builder
	for _, line := range aoc.InputLines() {
		for _, c := range iter.StringIterator(line, 1).List() {
			util.VecAdd(pos, motions[c])
			util.VecClamp(pos, 1)
			// part 2
			mot := motions[c]
			util.VecAdd(pos2, mot)
			v := g2.Get(pos2)
			if v == "*" || v == " " {
				util.VecAdd(pos2, util.VecInvert(mot))
			}
		}
		sb.WriteString(g.Get(pos))
		sb2.WriteString(g2.Get(pos2))
	}

	*p1 = sb.String()
	*p2 = sb2.String()
}
