package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"math"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2019, 3)
	aoc.Test(run, "sample1.txt", "6", "30")
	aoc.Test(run, "sample2.txt", "159", "610") // Is the sample wrong?
	aoc.Test(run, "sample3.txt", "135", "410")
	aoc.Test(run, "input.txt", "245", "")
	aoc.Run(run)
}

type wireset map[int]set.Set[int]

func run(p1 *string, p2 *string) {
	wires := make([]wireset, 2)
	wiredefs := aoc.InputLines()
	for idx, wiredef := range wiredefs {
		wire := make(map[int]set.Set[int])
		wires[idx] = wire
		trace(wiredef, func(x int, y int, stop *bool) {
			col, ok := wire[x]
			if !ok {
				wire[x] = set.New[int]()
				col = wire[x]
			}
			col.Insert(y)
		})
	}
	cx := crossovers(wires[0], wires[1])
	dists := f8l.Map(cx, func(c grid.Cell) int {
		return c[0] + c[1]
	})
	minDist := f8l.Reduce(dists, math.MaxInt32, util.IntMin)
	*p1 = strconv.Itoa(minDist)

	sigDelay := func(c grid.Cell) int {
		return wireDist(wiredefs[0], c) + wireDist(wiredefs[1], c)
	}

	delays := iter.Map[grid.Cell, int](iter.ListIterator(cx), sigDelay).List()
	minDelay := f8l.Reduce(delays, math.MaxInt32, util.IntMin)
	*p2 = strconv.Itoa(minDelay)
}

func trace(wiredef string, trav func(int, int, *bool)) {
	pos := [2]int{0, 0}
	stop := false
	steps := 0
	for _, leg := range strings.Split(wiredef, ",") {
		dir := rune(leg[0])
		mag := util.UnsafeAtoi(leg[1:])
		vec := map[rune][2]int{
			'L': [2]int{0, -1},
			'R': [2]int{0, 1},
			'U': [2]int{1, 0},
			'D': [2]int{-1, 0},
		}[dir]
		for i := 0; i < mag; i++ {
			steps++
			pos[0] += vec[0]
			pos[1] += vec[1]
			trav(pos[0], pos[1], &stop)
			if stop {
				return
			}
		}
	}
}

func crossovers(a wireset, b wireset) []grid.Cell {
	touches := make([]grid.Cell, 0)
	for x, _ := range a {
		for _, y := range set.Intersection(a[x], b[x]).Items() {
			touches = append(touches, grid.Cell{x, y})
		}
	}
	return touches
}

func wireDist(wiredef string, cell grid.Cell) int {
	dist := 1
	tt := make([]grid.Cell, 0)
	tt = append(tt, grid.Cell{0, 0})
	trace(wiredef, func(x int, y int, stop *bool) {
		tt = append(tt, grid.Cell{x, y})
		if cell[0] == x && cell[1] == y {
			*stop = true
		} else {
			dist++
		}
	})
	return dist
}
