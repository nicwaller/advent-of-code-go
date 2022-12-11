package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2016, 1)
	aoc.Test(run, "input.txt", "252", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	dirs := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	dirV := 0

	pos := grid.Cell{0, 0}
	seen := set.New[string]()
	for _, instruction := range strings.Split(aoc.InputString(), ", ") {
		switch instruction[0:1] {
		case "L":
			dirV = (dirV + 4 - 1) % 4
		case "R":
			dirV = (dirV + 4 + 1) % 4
		default:
			panic(instruction)
		}
		mag := util.UnsafeAtoi(instruction[1:])
		for i := 0; i < mag; i++ {
			util.VecAdd(pos, dirs[dirV])
			if *p2 == "" {
				cName := fmt.Sprintf("%d,%d", pos[0], pos[1])
				if seen.Contains(cName) {
					*p2 = strconv.Itoa(grid.ManhattanDistance(pos, grid.Cell{0, 0}))
				} else {
					seen.Add(cName)
				}
			}
		}
	}
	*p1 = strconv.Itoa(grid.ManhattanDistance(pos, grid.Cell{0, 0}))
}
