package main

import (
	"advent-of-code/lib/grid"
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
	content2 := parseFile()
	fmt.Printf("Part 2: %d\n", part2(content2))
}

type fileType struct {
	draws  []int
	boards []grid.Grid[int]
}

func parseFile() fileType {
	fbytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	//Split by line
	scanner := bufio.NewScanner(bytes.NewReader(fbytes))

	//Take 1 for Number stream, split by comma, parseInt.
	// numbers drawn
	scanner.Scan()
	drawNums := fmap[string, int](strings.Split(scanner.Text(), ","), unsafeAtoi)

	// bingo boards
	//Divide remainder into chunks of 6 lines
	var boards []grid.Grid[int]
	for {
		boardLines := scanMulti(scanner, 6)
		if len(boardLines) == 0 {
			break
		}
		//Prune the first (empty) line
		boardLines = boardLines[1:]
		boardStr := strings.Join(boardLines, "\n")
		board := grid.FromDelimitedStringAsInt(boardStr, ' ')
		boards = append(boards, board)
	}
	return fileType{
		draws:  drawNums,
		boards: boards,
	}
}

func scanMulti(scanner *bufio.Scanner, count int) []string {
	results := make([]string, count)
	for i := 0; i < count; i++ {
		if !scanner.Scan() {
			return []string{}
		}
		results[i] = scanner.Text()
		//results = append(results, scanner.Text())
	}
	return results
}

func unsafeAtoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func sum(values []int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func checkBingo(g grid.Grid[int]) bool {
	nextRow := g.RowIterator()
	for row := nextRow(); row != nil; row = nextRow() {
		if sum(row) == 500 {
			return true
		}
	}
	nextCol := g.ColumnIterator()
	for col := nextCol(); col != nil; col = nextCol() {
		if sum(col) == 500 {
			return true
		}
	}
	return false
}

func boardSum(g grid.Grid[int]) int {
	filterFn := func(v int) bool {
		return v != 100
	}
	remainingValues := ffilter[int](g.Values(), filterFn)
	return sum(remainingValues)
}

func part1(input fileType) int {
	for _, drawn := range input.draws {
		for _, board := range input.boards {
			if board.Replace(drawn, 100) > 0 {
				if checkBingo(board) {
					finalScore := drawn * boardSum(board)
					return finalScore // 16716
				}
			}
		}
	}
	return -1
}

func part2(input fileType) int {
	var lastWinScore int
	alreadyBingo := make([]bool, len(input.boards))
	for drawNumber, drawn := range input.draws {
		_ = drawNumber
		for boardId, board := range input.boards {
			if alreadyBingo[boardId] {
				continue
			}
			updates := board.Replace(drawn, 100)
			changed := updates > 0
			if changed {
				if checkBingo(board) {
					//fmt.Printf("got a bingo! on draw %d\n", drawNumber)
					lastWinScore = drawn * boardSum(board)
					alreadyBingo[boardId] = true
				}
			}
		}
	}
	if lastWinScore != 4880 {
		fmt.Println(lastWinScore)
		panic(4880)
	}
	return lastWinScore // 4880
}

func fmap[I interface{}, O interface{}](items []I, mapFn func(I) O) []O {
	var results []O = make([]O, len(items))
	for i, item := range items {
		results[i] = mapFn(item)
		//results = append(results, mapFn(item))
	}
	return results
}

func ffilter[T comparable](items []T, include func(T) bool) []T {
	var results = make([]T, 0)
	for _, item := range items {
		if include(item) {
			results = append(results, item)
		}
	}
	return results
}
