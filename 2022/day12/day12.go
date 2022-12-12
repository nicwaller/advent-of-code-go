package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"fmt"
	"github.com/yourbasic/graph"
	"strconv"
)

func main() {
	aoc.Select(2022, 12)
	aoc.Test(run, "sample.txt", "31", "")
	//aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	g := grid.FromString(aoc.InputString())
	gra := graph.New(g.Width() * g.Height())
	g.Print()
	allCells := g.Cells()
	for allCells.Next() {
		cell := allCells.Value()
		cellVal := g.Get(cell)
		cVN := int(cellVal[0])
		if cellVal == "S" {
			cVN = int('a')
		}
		offset := g.OffsetFromCell(cell)
		for _, ne := range g.NeighboursAdjacent(cell, false) {
			neVal := g.Get(ne)
			nVN := int(neVal[0])
			if neVal == "E" {
				nVN = 'z'
			}
			//jump := util.IntAbs(cVN - nVN)
			diff := nVN - cVN
			//if jump <= 1 {
			if diff <= 1 {
				offsetNe := g.OffsetFromCell(ne)
				gra.AddCost(offset, offsetNe, 1)
			}
		}
	}
	startCell := g.Filter(func(c grid.Cell, s string) bool { return s == "S" }).TakeFirst()
	goalCell := g.Filter(func(c grid.Cell, s string) bool { return s == "E" }).TakeFirst()
	startOffset := g.OffsetFromCell(startCell)
	goalOffset := g.OffsetFromCell(goalCell)
	p, dist := graph.ShortestPath(gra, startOffset, goalOffset)
	fmt.Println(p, dist)
	*p1 = strconv.Itoa(int(dist))

	a := analyze.Box[int]{}
	candidates := g.Filter(func(c grid.Cell, s string) bool { return s == "a" })
	for candidates.Next() {
		startOffset = g.OffsetFromCell(candidates.Value())
		_, dist = graph.ShortestPath(gra, startOffset, goalOffset)
		if dist != -1 {

			a.Put(int(dist))
		}
	}
	*p2 = strconv.Itoa(a.Min)
}
