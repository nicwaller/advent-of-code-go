package main

import (
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

type fileType struct {
	draws  []int
	boards []grid.Grid[int]
}

func parseFile() fileType {
	lines := util.ReadLines("input.txt")
	drawStr, lines := util.TakeOne(lines)
	drawStrs := strings.Split(drawStr, ",")
	drawNums := f8l.Atoi(&drawStrs)
	var boards []grid.Grid[int]
	for _, boardLines := range util.Chunk[string](lines, 6) {
		boardStr := strings.Join(boardLines[1:], "\n")
		board := grid.FromDelimitedStringAsInt(boardStr, ' ')
		boards = append(boards, board)
	}
	return fileType{
		draws:  drawNums,
		boards: boards,
	}
}

func checkBingo(g grid.Grid[int]) bool {
	var line *[]int
	next := iter.Chain(g.RowIteratorIter(), g.ColIteratorIter())
	for next(&line) {
		if f8l.Sum(line) == 500 {
			return true
		}
	}
	return false
}

func boardSum(g grid.Grid[int]) int {
	filterFn := func(v int) bool {
		return v != 100 // 100 is a special sentinel number in this file
	}
	remainingValues := f8l.Filter[int](g.Values(), filterFn)
	return f8l.Sum(&remainingValues)
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
