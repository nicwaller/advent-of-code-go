package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"strconv"
)

var gs grid.Slice

func main() {
	aoc.Select(2022, 14)
	gs = grid.SliceEnclosing(
		grid.Cell{0, 500},
		grid.Cell{9, 494},
		grid.Cell{9, 503},
	)
	aoc.Test(run, "sample.txt", "24", "93")
	gs = grid.SliceEnclosing(
		grid.Cell{0, 500},
		grid.Cell{300, 100},
		grid.Cell{300, 900},
	)
	//aoc.Test(run, "input.txt", "808", "26625")
	aoc.Run(run)
}

var maxY int

func parse() grid.Grid[string] {
	g := grid.NewGridFromSlice[string](gs)
	g.Fill(".")
	for _, line := range aoc.InputLines() {
		nf := util.NumberFields(line)
		for i := 0; i < len(nf)-2; i += 2 {
			c0 := grid.Cell{nf[i+1], nf[i+0]}
			c1 := grid.Cell{nf[i+3], nf[i+2]}
			for _, c := range grid.Line(c0, c1).List() {
				maxY = util.IntMax(maxY, c[0])
				g.Set(c, "#")
			}
		}
	}
	return g
}

func run(p1 *string, p2 *string) {
	g := parse()
	origin := grid.Cell{0, 500}
	restingPlace := func(start grid.Cell) grid.Cell {
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
				return p
			}
		}
		// abyss!
		return origin
	}
	for grains := 0; ; grains++ {
		rp := restingPlace(origin)
		g.Set(rp, "o")
		//g.Print()
		if rp[0] == origin[0] && rp[1] == origin[1] {
			// abyss!
			*p1 = strconv.Itoa(grains)
			break
		}
	}

	g = parse()
	// floor!
	g.Grow(10, ".")
	left := grid.Cell{0, -1}
	right := grid.Cell{0, 1}
	floorCentre := grid.Cell{maxY + 2, 500}
	g.Set(floorCentre, "#")
	for rayLeft := g.Ray(floorCentre, left); rayLeft.Next(); {
		g.Set(rayLeft.Value(), "#")
	}
	for rayRight := g.Ray(floorCentre, right); rayRight.Next(); {
		g.Set(rayRight.Value(), "#")
	}
	for grains := 0; ; {
		grains++
		rp := restingPlace(origin)
		g.Set(rp, "o")
		//g.Print()
		if rp[0] == origin[0] && rp[1] == origin[1] {
			// abyss!
			*p2 = strconv.Itoa(grains)
			break
		}
	}
	//g.Print()
}
