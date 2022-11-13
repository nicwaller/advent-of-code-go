package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
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
	fmt.Println(input)
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

func part2(g fileType) int {
	//assert.EqualAny(basinMultiplyResult, []int{1134, 1023660}, "basinMultiplyResult")
	return -1
}
