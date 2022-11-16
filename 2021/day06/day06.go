package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2021, 6)
	aoc.Test(run, "sample.txt", "5934", "26984457539")
	aoc.Test(run, "input.txt", "350605", "1592778185024")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	fish := make([]int, 9)
	for _, v := range util.NumberFields(aoc.InputString()) {
		fish[v]++
	}
	for day := 1; day <= 256; day++ {
		newFish := fish[0]
		for i := 0; i < 8; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += newFish
		fish[8] = newFish
		if day == 80 {
			*p1 = strconv.Itoa(f8l.Sum(fish))
		}
	}
	*p2 = strconv.Itoa(f8l.Sum(fish))
}
