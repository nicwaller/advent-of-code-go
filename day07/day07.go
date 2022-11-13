package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"fmt"
	"math"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

type fileType []int

func parseFile() []int {
	return util.NumberFields(aoc.InputString())
}

func part1(input []int) int {
	fuelRequired := func(hPos int) int {
		fuel := 0
		for _, crabX := range input {
			// TODO: move IntAbs out of grid library
			fuel += grid.IntAbs(crabX - hPos)
		}
		return fuel
	}
	globalMin := f8l.Reduce(&input, math.MaxInt32, util.IntMin)
	globalMax := f8l.Reduce(&input, math.MinInt32, util.IntMax)
	_ = fuelRequired(4)
	//assert.NotEqual(totalFish, 1711)

	// PERF: this could be much, MUCH more efficient.
	best := iter.Range(globalMin, globalMax).Map(fuelRequired).Reduce(util.IntMin, math.MaxInt32)
	//assert.NotEqual(totalFish, 5934)
	assert.EqualAny(best, []int{37, 345035}, "best")
	return best
}

func part2(input []int) int {
	moveCosts := make([]int, 2000)
	recalcMoveCosts := func() {
		s := 0
		for i, _ := range moveCosts {
			s += i
			moveCosts[i] = s
		}
	}
	recalcMoveCosts()
	highestDiff := 0
	fuelRequired := func(hPos int) int {
		fuel := 0
		for _, crabX := range input {
			// TODO: move IntAbs out of grid library
			diff := grid.IntAbs(crabX - hPos)
			if diff > highestDiff {
				highestDiff = diff
				fmt.Println(diff)
			}
			fuel += moveCosts[diff]
		}
		return fuel
	}
	globalMin := f8l.Reduce(&input, math.MaxInt32, util.IntMin)
	globalMax := f8l.Reduce(&input, math.MinInt32, util.IntMax)
	_ = fuelRequired(4)
	//assert.NotEqual(totalFish, 1711)

	// PERF: this could be much, MUCH more efficient.
	best := iter.Range(globalMin, globalMax).Map(fuelRequired).Reduce(util.IntMin, math.MaxInt32)
	//assert.NotEqual(totalFish, 5934)
	assert.EqualAny(best, []int{168, 97038163}, "best")
	return best
}
