package main

import (
	"fmt"
	"strconv"

	"advent-of-code/lib/aoc"
)

func main() {
	aoc.Select(2025, 4)
	aoc.Test(run, "sample.txt", "13", "1395")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	accessible := 0
	g := aoc.InputGridRunes()
	for _, cell := range g.Cells().List() {
		if g.Get(cell) == "." {
			continue
		}

		count := 0
		for _, n := range g.NeighboursSurround(cell, false) {
			if g.Get(n) == "@" {
				count++
			}
		}
		if count < 4 {
			fmt.Println(cell)
			accessible++
		}
	}
	g.Print()

	//for _, line := range aoc.InputLines() {
	//}
	*p1 = strconv.Itoa(accessible)
	*p2 = ""
}
