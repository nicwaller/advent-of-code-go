package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"fmt"
	"github.com/yourbasic/graph"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

func parseFile() grid.Grid[int] {
	return grid.FromStringAsInt(aoc.InputString())
}

func part1(input grid.Grid[int]) int {
	g := graph.New(input.RowCount() * input.ColumnCount())
	for c := input.Cells(); c.Next(); {
		dst := input.OffsetFromCell(c.Value())
		for _, n := range input.NeighboursAdjacent(c.Value(), false) {
			src := input.OffsetFromCell(n)
			g.AddCost(src, dst, int64(input.Get(c.Value())))
		}
	}
	_, distance := graph.ShortestPath(g, 0, input.RowCount()*input.ColumnCount()-1)
	return int(distance)
}

func part2(inputA grid.Grid[int]) int {
	bigGrid := grid.NewGrid[int](inputA.RowCount()*5, inputA.ColumnCount()*5)
	for dstCellIter := bigGrid.Cells(); dstCellIter.Next(); {
		dstCell := dstCellIter.Value()
		srcCell := grid.Cell{
			dstCell[0] % inputA.Width(),
			dstCell[1] % inputA.Height(),
		}
		riskOffset := dstCell[0]/inputA.Width() + dstCell[1]/inputA.Height()
		srcVal := inputA.Get(srcCell)
		// this line here is a bit of straight fuckery
		// y u do this to me
		// > "However, risk levels above 9 wrap back around to 1"
		// SERIOUSLY?!
		dstVal := (srcVal+riskOffset)%10 + ((srcVal + riskOffset) / 10)
		_ = riskOffset
		bigGrid.Set(dstCell, dstVal)
	}
	bigGrid.Print()
	g := graph.New(bigGrid.RowCount() * bigGrid.ColumnCount())
	for c := bigGrid.Cells(); c.Next(); {
		dst := bigGrid.OffsetFromCell(c.Value())
		for _, n := range bigGrid.NeighboursAdjacent(c.Value(), false) {
			src := bigGrid.OffsetFromCell(n)
			g.AddCost(src, dst, int64(bigGrid.Get(c.Value())))
		}
	}
	_, distance := graph.ShortestPath(g, 0, bigGrid.RowCount()*bigGrid.ColumnCount()-1)
	return int(distance)
}
