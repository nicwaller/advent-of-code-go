package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"fmt"
	"math"
	"strings"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

type fileType struct {
	polymerTemplate string
	rules           map[string]string
}

func parseFile() fileType {
	ft := fileType{
		polymerTemplate: "",
		rules:           make(map[string]string),
	}
	lines := aoc.InputLinesIterator()
	ft.polymerTemplate = lines.TakeFirst()
	_ = lines.Skip(1)

	for lines.Next() {
		parts := strings.Split(lines.Value(), " -> ")
		k, v := parts[0], parts[1]
		ft.rules[k] = v
	}

	return ft
}

func polymerize(template string, rules map[string]string) string {
	var sb strings.Builder
	foo := iter.SlidingWindow[string](2, iter.StringIterator(template, 1))
	for foo.Next() {
		pair := foo.Value()
		sb.WriteString(pair[0])
		sb.WriteString(rules[strings.Join(pair, "")])
	}
	sb.WriteString(foo.Value()[1])
	return sb.String()
}

func part1(input fileType) int {
	//fmt.Println(input)
	cur := input.polymerTemplate
	for step := 1; step <= 10; step++ {
		cur = polymerize(cur, input.rules)
	}
	stat := analyze.CountDistinct(iter.StringIterator(cur, 1).List())
	// TODO: extend the analyze package for this
	most := math.MinInt32
	least := math.MaxInt32
	for _, v := range stat {
		most = util.IntMax(most, v)
		least = util.IntMin(least, v)
	}
	return most - least
}

func part2(input fileType) int {
	//fmt.Println(input)

	// prepare the empty map of pairs
	pairs := make(map[string]int)
	for pair, _ := range input.rules {
		pairs[pair] = 0
	}

	// re-interpret the template as a set of pairs
	for i := 0; i < len(input.polymerTemplate)-1; i++ {
		pair := input.polymerTemplate[i : i+2]
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
			middlechar := input.rules[pair]
			newPairs[pair] -= count
			leftPair := string(pair[0]) + middlechar
			rightPair := middlechar + string(pair[1])
			newPairs[leftPair] += count
			newPairs[rightPair] += count
		}

		return newPairs
	}

	//fmt.Println(pairs)
	for step := 1; step <= 40; step++ {
		pairs = iterate(pairs)
		//fmt.Println(pairs)
	}

	aggr := func(m map[string]int) map[string]int {
		mm := make(map[string]int)
		for k, v := range m {
			x := string(k[0])
			mm[x] += v
		}
		lastChar := input.polymerTemplate[len(input.polymerTemplate)-1]
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
	diff := most - least
	fmt.Printf("most: %v, least: %v, diff: %v\n", most, least, diff)
	assert.NotEqual(diff, 2745)
	assert.NotEqual(diff, 4397257877492)
	assert.EqualAny(diff, []int{2188189693529, 3420801168962}, "diff")
	return most - least
}
