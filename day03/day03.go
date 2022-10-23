package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := parseFile()
	fmt.Printf("Part 1: %d\n", part1(content))
	fmt.Printf("Part 2: %d\n", part2(content))
}

type fileType [][]uint8

func parseFile() fileType {
	fbytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var rows [][]uint8
	scanner := bufio.NewScanner(bytes.NewReader(fbytes))
	for scanner.Scan() {
		row := make([]uint8, 12)
		for index, char := range scanner.Text() {
			val, _ := strconv.Atoi(string(char))
			row[index] = uint8(val)
		}
		rows = append(rows, row)
	}
	return rows
}

func transpose[T comparable](matrix [][]T) [][]T {
	oldSizeInner := len(matrix[0])
	oldSizeOuter := len(matrix)
	newSizeInner := oldSizeOuter
	newSizeOuter := oldSizeInner
	newRows := make([][]T, newSizeOuter)
	for i := 0; i < newSizeOuter; i++ {
		newRows[i] = make([]T, newSizeInner)
	}
	fmt.Printf("transposing into %d by %d\n", newSizeOuter, newSizeInner)
	for i := 0; i < newSizeOuter; i++ {
		for j := 0; j < newSizeInner; j++ {
			newRows[i][j] = matrix[j][i]
		}
	}
	return newRows
}

func gammaRate(input fileType) int {
	var inputX [][]uint8
	inputX = input
	tResult := transpose[uint8](inputX)
	gammaRate := 0
	for colIndex, column := range tResult {
		_ = colIndex
		distinct := countDistinct(column)
		zeroCount := distinct[0]
		oneCount := distinct[1]
		shiftPosition := 11 - colIndex
		if oneCount > zeroCount {
			gammaRate += 1 << shiftPosition
		}
	}
}

func part1(input fileType) int {
	gammaRate := gammaRate(input)
	epsilonRate := gammaRate ^ 0b0000111111111111
	return gammaRate * epsilonRate
}

func part2(input fileType) int {
	var numbers = make([]int64, len(input))
	for i := 0; i < len(input); i++ {
		intArray := input[i]
		strArray := fmap[uint8, string](intArray, func(i uint8) string {
			return strconv.Itoa(int(i))
		})
		strJoined := strings.Join(strArray, "")
		numericValue, err := strconv.ParseInt(strJoined, 2, 16)
		if err != nil {
			fmt.Println(err)
		}
		numbers[i] = numericValue
	}

	var keep = make([]bool, len(numbers))
	for i := 0; i < len(keep); i++ {
		keep[i] = true
	}
	kept := len(numbers)

	// counting bits from left to right, like a heathen
	for bitPosition := 0; bitPosition < 12; bitPosition++ {
		for i := 0; i < len(keep); i++ {
			if !keep[i] {
				continue
			}
			// FML I hate this one
		}
		fmt.Println(kept)
	}
	return 0
}

func countDistinct[T uint8](list []T) map[T]int {
	keys := make(map[T]int)
	for _, item := range list {
		//fmt.Println(item)
		if _, ok := keys[item]; !ok {
			keys[item] = 0
		}
		keys[item] += 1
	}
	return keys
}

func fmap[I interface{}, O interface{}](items []I, mapFn func(I) O) []O {
	var results []O = make([]O, len(items))
	for _, item := range items {
		results = append(results, mapFn(item))
	}
	return results
}
