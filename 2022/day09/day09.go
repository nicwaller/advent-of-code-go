package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 9)
	aoc.Test(run, "sample.txt", "13", "1")
	aoc.Test(run, "sample2.txt", "", "36")
	aoc.Test(run, "input.txt", "6470", "2658")
	aoc.Out()
}

func tailChase(head *grid.Cell, tail *grid.Cell) {
	switch grid.ChebyshevDistance(*head, *tail) {
	case 0:
		// rope segments may overlap, especially at the beginning
	case 1:
		// adjacent segments don't need to move
	case 2:
		v := util.VecDiff(*head, *tail)
		util.VecClamp(v, 1)
		util.VecAdd(*tail, v)
	}
}

func ropesim(rope []grid.Cell) {
	for i := 1; i < len(rope); i++ {
		tailChase(&rope[i-1], &rope[i])
	}
}

func run(p1 *string, p2 *string) {
	dirMap := map[string]grid.Cell{
		"U": {1, 0},
		"D": {-1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	shortrope := util.Make[grid.Cell](2, func() grid.Cell { return grid.Cell{0, 0} })
	longrope := util.Make[grid.Cell](10, func() grid.Cell { return grid.Cell{0, 0} })
	seen1 := set.New[int]()
	seen2 := set.New[int]()
	for _, line := range aoc.InputLines() {
		moveDirection := dirMap[strings.Fields(line)[0]]
		moveDistance := util.UnsafeAtoi(strings.Fields(line)[1])
		for step := 1; step <= moveDistance; step++ {
			util.VecAdd(shortrope[0], moveDirection)
			util.VecAdd(longrope[0], moveDirection)
			ropesim(shortrope)
			ropesim(longrope)
			seen1.Add(grid.CellHash(shortrope[1]))
			seen2.Add(grid.CellHash(longrope[9]))
		}
	}
	*p1 = strconv.Itoa(seen1.Size())
	*p2 = strconv.Itoa(seen2.Size())
}
