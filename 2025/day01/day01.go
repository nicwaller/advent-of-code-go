package main

import (
	"advent-of-code/lib/aoc"
)

func main() {
	aoc.Select(2025, 1)
	aoc.Test(run, "sample.txt", "142", "")
	aoc.Test(run, "sample2.txt", "", "281")
	aoc.Test(run, "input.txt", "55090", "54845")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	//for _, line := range aoc.InputLines() {
	//}

	*p1 = ""
	*p2 = ""
}
