package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"math"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2021, 14)
	aoc.Test(run, "sample.txt", "1588", "2188189693529")
	aoc.Test(run, "input.txt", "2745", "3420801168962")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLinesIterator()
	polymerTemplate := lines.MustTakeArray(2)[0]

	rules := make(map[string]string)
	for lines.Next() {
		parts := strings.Split(lines.Value(), " -> ")
		rules[parts[0]] = parts[1]
	}

	// prepare the empty map of pairs
	pairs := make(map[string]int)
	for pair, _ := range rules {
		pairs[pair] = 0
	}

	// re-interpret the template as a set of pairs
	for i := 0; i < len(polymerTemplate)-1; i++ {
		pair := polymerTemplate[i : i+2]
		pairs[pair]++
	}

	iterate := func(curPairs map[string]int) map[string]int {
		// make a deep copy
		newPairs := make(map[string]int)
		for pair, count := range curPairs {
			newPairs[pair] = count
		}

		// populate the next generation
		for pair, count := range curPairs {
			middlechar := rules[pair]
			newPairs[pair] -= count
			leftPair := string(pair[0]) + middlechar
			rightPair := middlechar + string(pair[1])
			newPairs[leftPair] += count
			newPairs[rightPair] += count
		}

		return newPairs
	}

	mostLeast := func() (int, int) {
		aggr := func(m map[string]int) map[string]int {
			mm := make(map[string]int)
			for k, v := range m {
				x := string(k[0])
				mm[x] += v
			}
			lastChar := polymerTemplate[len(polymerTemplate)-1]
			mm[string(lastChar)]++
			return mm
		}
		stat := aggr(pairs)
		//fmt.Println(stat)

		most := math.MinInt64
		least := math.MaxInt64
		for _, v := range stat {
			most = util.IntMax(most, v)
			least = util.IntMin(least, v)
		}
		return most, least
	}

	//fmt.Println(pairs)
	for step := 1; step <= 10; step++ {
		pairs = iterate(pairs)
	}

	most1, least1 := mostLeast()
	diff1 := most1 - least1
	*p1 = strconv.Itoa(diff1)

	for step := 11; step <= 40; step++ {
		pairs = iterate(pairs)
	}

	most2, least2 := mostLeast()
	diff2 := most2 - least2
	*p2 = strconv.Itoa(diff2)
}
