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
	p, d := graph.ShortestPath(g, 0, input.RowCount()*input.ColumnCount()-1)
	fmt.Println(p)
	fmt.Println(d)
	pathCost := 0
	for _, step := range p[1:] {
		v := input.Get(input.CellFromOffset(step))
		fmt.Println(v)
		pathCost += v
	}
	return pathCost
}

func part2(input grid.Grid[int]) int {
	return -1
}
