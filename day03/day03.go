package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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

func part1(input fileType) int {
	fmt.Printf("input: %d by %d\n", len(input), len(input[0]))
	//x := countDistinct(input[0])
	var inputX [][]uint8
	inputX = input
	tResult := transpose[uint8](inputX)
	fmt.Printf("tResult: %d by %d\n", len(tResult), len(tResult[0]))
	gammaRate := 0
	for colIndex, column := range tResult {
		_ = colIndex
		distinct := countDistinct(column)
		zeroCount := distinct[0]
		oneCount := distinct[1]
		fmt.Println(zeroCount, oneCount)
		shiftPosition := 11 - colIndex
		if oneCount > zeroCount {
			fmt.Println("enabling bit in position %d", shiftPosition)
			gammaRate += 1 << shiftPosition
		}
	}
	epsilonRate := gammaRate ^ 0b0000111111111111
	return gammaRate * epsilonRate
}

func part2(input fileType) int {
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
