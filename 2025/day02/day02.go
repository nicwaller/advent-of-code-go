package main

import (
	"advent-of-code/lib/aoc"
)

func main() {
	aoc.Select(2025, 2)
	aoc.Test(run, "sample.txt", "", "")
	//aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	for _, line := range aoc.InputLines() {
		_ = line
	}
}
