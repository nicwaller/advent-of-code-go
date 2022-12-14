package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/stack"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2022, 14)
	aoc.Test(run, "sample.txt", "24", "93")
	aoc.Test(run, "input.txt", "808", "26625")
	aoc.Run(run)
	aoc.Out()
}

var maxY int

func parse(includeFloor bool) grid.Grid[string] {
	lineSegments := make([][]grid.Cell, 0)
	scanX := analyze.Box[int]{}
	scanY := analyze.Box[int]{}
	for _, line := range aoc.InputLines() {
		nf := util.NumberFields(line)
		for i := 0; i < len(nf); i += 2 {
			scanX.Put(nf[i])
			scanY.Put(nf[i+1])
		}
		for i := 0; i < len(nf)-2; i += 2 {
			c0 := grid.Cell{nf[i+1], nf[i+0]}
			c1 := grid.Cell{nf[i+3], nf[i+2]}
			lineSegments = append(lineSegments, []grid.Cell{c0, c1})
		}
	}
	if includeFloor {
		scanY.Put(scanY.Max + 2)
	}
	g := grid.NewGridFromSlice[string](grid.SliceEnclosing(
		grid.Cell{0, 500}, // sand origin
		grid.Cell{scanY.Max, scanX.Min - 1 - scanY.Max},
		grid.Cell{scanY.Max, scanX.Max + 1 + scanY.Max},
	))
	g.Fill(".")
	for _, lineSegment := range lineSegments {
		for _, c := range grid.Line(util.Pair(lineSegment)).List() {
			maxY = util.IntMax(maxY, c[0])
			g.Set(c, "#")
		}
	}
	if includeFloor {
		floorCentre := grid.Cell{scanY.Max, 500}
		g.Set(floorCentre, "#")
		iter.Chain(
			g.Ray(floorCentre, grid.Cell{0, -1}),
			g.Ray(floorCentre, grid.Cell{0, 1}),
		).Each(func(c grid.Cell) {
			g.Set(c, "#")
		})
	}
	return g
}

func run(p1 *string, p2 *string) {
	g := parse(false)
	origin := grid.Cell{0, 500}
	tk := stack.NewStack[grid.Cell]()
	restingPlace := func(start grid.Cell) grid.Cell {
		if tk.Height() > 0 {
			start = tk.Peek()
		}
		p := util.Copy(start)
		for p[0] < g.Height()-1 {
			if g.Get(grid.Cell{p[0] + 1, p[1]}) == "." {
				// down is good
				p[0]++
			} else if g.Get(grid.Cell{p[0] + 1, p[1] - 1}) == "." {
				// down-left is good too
				p[0]++
				p[1]--
			} else if g.Get(grid.Cell{p[0] + 1, p[1] + 1}) == "." {
				// down-right is good too
				p[0]++
				p[1]++
			} else {
				_, _ = tk.Pop()
				return p
			}
			tk.Push(util.Copy(p))
		}
		// abyss!
		return origin
	}

	fillGrains := func() int {
		for grains := 0; ; grains++ {
			rp := restingPlace(origin)
			g.Set(rp, "o")
			if rp[0] == origin[0] && rp[1] == origin[1] {
				return grains
			}
		}
	}

	*p1 = strconv.Itoa(fillGrains())
	g = parse(true)
	*p2 = strconv.Itoa(fillGrains() + 1)
}
