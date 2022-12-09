package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 9)
	aoc.Test(run, "sample.txt", "13", "1")
	aoc.Test(run, "sample2.txt", "", "36")
	aoc.Test(run, "input.txt", "6470", "2658")
	aoc.Run(run)
	aoc.Out()
}

func manhattanDistance(a grid.Cell, b grid.Cell) int {
	return util.IntAbs(a[0]-b[0]) + util.IntAbs(a[1]-b[1])
}

func run(p1 *string, p2 *string) {
	head := grid.Cell{0, 0}
	tail := grid.Cell{0, 0}

	g := grid.NewGrid[int](1, 1)
	g.Grow(300, 0)

	tailChase := func(head *grid.Cell, tail *grid.Cell) {
		if (*tail)[0] == (*head)[0] {
			d := (*head)[1] - (*tail)[1]
			if util.IntAbs(d) == 2 {
				(*tail)[1] += d / 2
			}
		} else if (*tail)[1] == (*head)[1] {
			d := (*head)[0] - (*tail)[0]
			if util.IntAbs(d) == 2 {
				(*tail)[0] += d / 2
			}
		} else if (*tail)[0] != (*head)[0] && (*tail)[1] != (*head)[1] && manhattanDistance(*head, *tail) > 2 {
			// move diagonal closer
			if (*tail)[0] < (*head)[0] {
				(*tail)[0]++
			}
			if (*tail)[0] > (*head)[0] {
				(*tail)[0]--
			}
			if (*tail)[1] < (*head)[1] {
				(*tail)[1]++
			}
			if (*tail)[1] > (*head)[1] {
				(*tail)[1]--
			}
		}

	}

	mm := map[string]grid.Cell{
		"U": {1, 0},
		"D": {-1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	for _, line := range aoc.InputLines() {
		dir, mag := util.Pair(strings.Fields(line))
		vec := mm[dir]
		for step := 1; step <= util.UnsafeAtoi(mag); step++ {
			for d := 0; d <= 1; d++ {
				head[d] += vec[d]
			}
			tailChase(&head, &tail)
			g.Set(tail, 1)
			//fmt.Printf("%v, %v\n", head, tail)
		}
	}
	*p1 = strconv.Itoa(g.Filter(func(c grid.Cell, v int) bool {
		return v > 0
	}).Count())

	g.Fill(0)
	longrope := make([]grid.Cell, 10)
	for i := 0; i < len(longrope); i++ {
		longrope[i] = make([]int, 2)
	}
	for _, line := range aoc.InputLines() {
		dir, mag := util.Pair(strings.Fields(line))
		vec := mm[dir]
		for step := 1; step <= util.UnsafeAtoi(mag); step++ {
			for d := 0; d <= 1; d++ {
				longrope[0][d] += vec[d]
			}
			for node := 1; node < len(longrope); node++ {
				tailChase(&longrope[node-1], &longrope[node])
			}
			g.Set(longrope[9], 1)
		}
		fmt.Printf("%v, %v\n", longrope[0], longrope[9])
	}

	*p2 = strconv.Itoa(g.Filter(func(c grid.Cell, v int) bool {
		return v > 0
	}).Count())

}
