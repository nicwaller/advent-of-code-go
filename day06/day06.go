package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"fmt"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

type fileType []int

func parseFile() fileType {
	return util.NumberFields(aoc.InputString())
}

func part1(input fileType) int {
	fish := make([]int, 9)
	for _, v := range input {
		fish[v]++
	}
	for day := 1; day <= 256; day++ {
		newFish := fish[0]
		for i := 0; i < 8; i++ {
			fish[i] = fish[i+1]
		}
		fish[6] += newFish
		fish[8] = newFish
	}
	totalFish := f8l.Sum(&fish)
	assert.NotEqual(totalFish, 1711)
	assert.NotEqual(totalFish, 5934)
	assert.EqualAny(totalFish, []int{26, 350605, 1592778185024}, "total fish")
	return totalFish
}

func part2(g fileType) int {
	//assert.EqualAny(basinMultiplyResult, []int{1134, 1023660}, "basinMultiplyResult")
	return -1
}
