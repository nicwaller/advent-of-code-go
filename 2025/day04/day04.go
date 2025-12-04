package main

import (
	"fmt"
	"strconv"

	"advent-of-code/lib/aoc"
)

func main() {
	aoc.Select(2025, 4)
	aoc.Test(run, "sample.txt", "13", "43")
	aoc.Test(run, "input.txt", "1395", "8451")
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
			accessible++
		}
	}

	removedTotal := 0
cleanup:
	for {
		removed := 0

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
				g.Set(cell, ".")
				removed++
			}
		}

		fmt.Printf("removed this round: %d\n", removed)
		removedTotal += removed
		if removed == 0 {
			break cleanup
		}
	}

	*p1 = strconv.Itoa(accessible)
	*p2 = strconv.Itoa(removedTotal)
}
