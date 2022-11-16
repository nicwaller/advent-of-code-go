package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"strconv"
)

func main() {
	aoc.Select(2020, 3)
	aoc.Test(run, "sample.txt", "7", "336")
	aoc.Test(run, "input.txt", "148", "727923200")
	aoc.Run(run)
}

func countCollisions(g grid.Grid[string], vx int, vy int) int {
	x := 0
	y := 0
	trees := 0
	for y < g.RowCount() {
		if g.Get([]int{y, x % g.ColumnCount()}) == "#" {
			trees++
		}
		x += vx
		y += vy
	}
	return trees
}

func run(p1 *string, p2 *string) {
	g := grid.FromString(aoc.InputString())
	*p1 = strconv.Itoa(countCollisions(g, 3, 1))

	*p2 = strconv.Itoa(1 *
		countCollisions(g, 1, 1) *
		countCollisions(g, 3, 1) *
		countCollisions(g, 5, 1) *
		countCollisions(g, 7, 1) *
		countCollisions(g, 1, 2))
}
