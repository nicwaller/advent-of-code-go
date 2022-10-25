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
	drawStrs := strings.Split(lines.TakeFirst(), ",")
	var boards []grid.Grid[int]
	for chunk := iter.Chunk(6, lines); chunk.Next(); {
		boardStr := strings.Join(chunk.Value(), "\n")
		boards = append(boards, grid.FromDelimitedStringAsInt(boardStr, ' '))
	}
	return fileType{
		draws:  f8l.Atoi(&drawStrs),
		boards: boards,
	}
}

func checkBingo(g *grid.Grid[int]) bool {
	for lines := iter.Chain(g.RowIter(), g.ColumnIter()); lines.Next(); {
		if f8l.Sum(lines.Value()) == 500 {
			return true
		}
	}
	return false
}

func boardSum(g *grid.Grid[int]) int {
	filterFn := func(v int) bool {
		return v != 100 // 100 is a special sentinel number in this file
	}
	sumFn := func(a int, b int) int {
		return a + b
	}
	return iter.ListIterator(g.Values()).Filter(filterFn).Reduce(sumFn, 0)
}

func part1(input fileType) int {
	for _, drawn := range input.draws {
		for _, board := range input.boards {
			if board.Replace(drawn, 100) > 0 && checkBingo(&board) {
				finalScore := drawn * boardSum(&board)
				return finalScore // 16716
			}
		}
	}
	return -1
}

func part2(input fileType) int {
	var lastWinScore int
	alreadyBingo := make([]bool, len(input.boards))
	for _, drawn := range input.draws {
		for boardId, board := range input.boards {
			if !alreadyBingo[boardId] && board.Replace(drawn, 100) > 0 && checkBingo(&board) {
				lastWinScore = drawn * boardSum(&board)
				alreadyBingo[boardId] = true
			}
		}
	}
	return lastWinScore // 4880
}
