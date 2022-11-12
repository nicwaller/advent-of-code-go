package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"fmt"
	"math"
	"sort"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

func parseFile() grid.Grid[int] {
	g := aoc.InputGridNumbers()
	return g
}

func part1(g grid.Grid[int]) int {
	riskSum := 0
	for iterCell := g.Cells(); iterCell.Next(); {
		cell := iterCell.Value()
		neighbours := g.NeighboursAdjacent(cell, false)
		iterNeighbours := iter.ListIterator(&neighbours)
		iterVals := iter.Transform[grid.Cell, int](iterNeighbours, func(c grid.Cell) int {
			return g.Get(c)
		})
		lowestNeighbour := iterVals.Reduce(util.IntMin, math.MaxInt32)
		cellValue := g.Get(cell)
		if cellValue < lowestNeighbour {
			fmt.Println(cell)
			riskLevel := 1 + cellValue
			riskSum += riskLevel
		}
	}
	assert.EqualAny(riskSum, []int{15, 516}, "riskSum")
	return riskSum
}

func part2(g grid.Grid[int]) int {
	basinSizes := make([]int, 0)
	for iterCell := g.Cells(); iterCell.Next(); {
		cell := iterCell.Value()
		if g.Get(cell) == 9 {
			continue
		}
		isBoundary := func(v int) bool { return v != 9 }
		changed := g.FloodFill(cell, isBoundary, 9)
		basinSizes = append(basinSizes, changed)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	basinMultiplyResult := basinSizes[0] * basinSizes[1] * basinSizes[2]
	assert.EqualAny(basinMultiplyResult, []int{1134, 1023660}, "basinMultiplyResult")
	return basinMultiplyResult
}
