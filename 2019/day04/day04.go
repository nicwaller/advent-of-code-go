package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2019, 4)
	aoc.Test(run, "input.txt", "1625", "1111")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	min, max := util.Pair(f8l.Map(strings.Split(aoc.InputString(), "-"), util.UnsafeAtoi))
	pows := []int{100000, 10000, 1000, 100, 10, 1}
	digits := make([]int, 6)
	groupSize := make([]int, 6)
	count1 := 0
	count2 := 0
	for num := min; num <= max; num++ {
		for pv := 0; pv < 6; pv++ {
			digits[pv] = num / pows[pv] % 10
		}
		// sliding window
		neverDecreasing := true
		foundAdjacent1 := false
		foundAdjacent2 := false
		seqDigit := 10
		seqCount := 1

		for i := 0; i < 6; {
			j := i
			for ; j < 5 && digits[j] == digits[j+1]; j++ {
				seqCount++
			}
			for k := j; k >= i; k-- {
				groupSize[k] = seqCount
			}
			i += seqCount
			seqCount = 1
		}

		for x := 1; x < 6; x++ {
			if digits[x] != seqDigit {
				seqCount++
			} else {
				seqCount = 1
			}
			foundAdjacent1 = foundAdjacent1 || (digits[x] == digits[x-1])
			foundAdjacent2 = foundAdjacent2 || (digits[x] == digits[x-1] && groupSize[x] <= 2)
			neverDecreasing = neverDecreasing && (digits[x] >= digits[x-1])
		}
		if foundAdjacent1 && neverDecreasing {
			count1++
		}
		if foundAdjacent2 && neverDecreasing {
			count2++
		}
	}
	*p1 = strconv.Itoa(count1)
	*p2 = strconv.Itoa(count2)
}
