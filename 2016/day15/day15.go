package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2016, 15)
	aoc.Test(run, "sample.txt", "5", "")
	aoc.Test(run, "input.txt", "317371", "2080951")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	discSize := make([]int, 0)
	discStartingPos := make([]int, 0)
	for _, line := range aoc.InputLines() {
		nf := util.NumberFields(line)
		discSize = append(discSize, nf[1])
		discStartingPos = append(discStartingPos, nf[3])
	}

	leastSafeDelay := func() int {
		discPos := make([]int, len(discStartingPos))
		rotate := func(n int) {
			for i, _ := range discPos {
				discPos[i] += n
				discPos[i] %= discSize[i]
			}
		}
	delayLoop:
		for delay := 0; ; delay++ {
			copy(discPos, discStartingPos)
			rotate(delay)
			// follow the capsule through the discs
			for step := 0; step < len(discPos); step++ {
				rotate(1)
				if discPos[step] != 0 {
					continue delayLoop
				}
			}
			return delay
		}
	}

	*p1 = strconv.Itoa(leastSafeDelay())
	discSize = append(discSize, 11)
	discStartingPos = append(discStartingPos, 0)
	*p2 = strconv.Itoa(leastSafeDelay())
}
