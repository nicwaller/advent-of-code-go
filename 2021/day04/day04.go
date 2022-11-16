package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2021, 4)
	aoc.Test(run, "sample.txt", "4512", "1924")
	aoc.Test(run, "input.txt", "16716", "4880")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLinesIterator()
	draws := util.NumberFields(lines.TakeFirst())
	var boards []grid.Grid[int]
	for chunk := iter.Chunk(6, lines); chunk.Next(); {
		boardStr := strings.Join(chunk.Value(), "\n")
		boards = append(boards, grid.FromDelimitedStringAsInt(boardStr, ' '))
	}

	alreadyBingo := make([]bool, len(boards))
	for _, drawn := range draws {
		for boardId, board := range boards {
			changed := board.Replace(drawn, 100)
			if alreadyBingo[boardId] || changed == 0 {
				continue
			}
			if checkBingo(&board) {
				finalScore := drawn * boardSum(board)
				alreadyBingo[boardId] = true
				if *p1 == "" {
					*p1 = strconv.Itoa(finalScore)
				}
				*p2 = strconv.Itoa(finalScore)
			}
		}
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

func boardSum(g grid.Grid[int]) int {
	filterFn := func(v int) bool { return v != 100 }
	sumFn := func(a int, b int) int { return a + b }
	return g.ValuesIterator().Filter(filterFn).Reduce(sumFn, 0)
}
