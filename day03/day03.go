package main

import (
	"advent-of-code/day03/analyze"
	"advent-of-code/day03/grid"
	"fmt"
	"os"
)

func main() {
	content := parseFile()
	fmt.Printf("Part 1: %d\n", part1(content))
	fmt.Printf("Part 2: %d\n", part2(content))
}

func parseFile() grid.Grid[int] {
	fbytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	return grid.FromStringAsInt(string(fbytes))
}

func gammaRate(g grid.Grid[int]) int {
	gammaRate := 0
	for c := 0; c < g.ColumnCount(); c++ {
		column := g.Column(c)
		if analyze.MostCommon(column) == 1 {
			shiftPosition := 11 - c
			gammaRate += 1 << shiftPosition
		}
	}
	return gammaRate
}

func part1(input grid.Grid[int]) int {
	gammaRate := gammaRate(input)
	epsilonRate := gammaRate ^ 0b0000111111111111
	return gammaRate * epsilonRate
}

func part2(input grid.Grid[int]) int {
	//var numbers = make([]int64, len(input))
	//for i := 0; i < len(input); i++ {
	//	intArray := input[i]
	//	strArray := fmap[uint8, string](intArray, func(i uint8) string {
	//		return strconv.Itoa(int(i))
	//	})
	//	strJoined := strings.Join(strArray, "")
	//	numericValue, err := strconv.ParseInt(strJoined, 2, 16)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	numbers[i] = numericValue
	//}
	//
	//var keep = make([]bool, len(numbers))
	//for i := 0; i < len(keep); i++ {
	//	keep[i] = true
	//}
	//kept := len(numbers)
	//
	//// counting bits from left to right, like a heathen
	//for bitPosition := 0; bitPosition < 12; bitPosition++ {
	//	for i := 0; i < len(keep); i++ {
	//		if !keep[i] {
	//			continue
	//		}
	//		// FML I hate this one
	//	}
	//	fmt.Println(kept)
	//}
	return 0
}

func fmap[I interface{}, O interface{}](items []I, mapFn func(I) O) []O {
	var results []O = make([]O, len(items))
	for _, item := range items {
		results = append(results, mapFn(item))
	}
	return results
}
