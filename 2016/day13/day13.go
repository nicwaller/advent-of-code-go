package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"github.com/yourbasic/graph"
	"math/bits"
	"strconv"
)

var gridSize int
var target grid.Cell

func main() {
	aoc.Select(2016, 13)

	gridSize = 10
	target = grid.Cell{4, 7}
	aoc.TestLiteral(run, "10", "11", "25")

	gridSize = 40
	target = grid.Cell{39, 31}
	aoc.Test(run, "input.txt", "82", "138")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	n := util.UnsafeAtoi(aoc.InputString())
	g := grid.NewGrid[string](gridSize, gridSize)
	g.FillFunc2D(func(v string, x int, y int) string {
		a := x*x + 3*x + 2*x*y + y + y*y + n
		isWall := bits.OnesCount(uint(a))%2 == 1
		return map[bool]string{false: ".", true: "#"}[isWall]
	})

	gra := graph.New(g.Width() * g.Height())
	openCells := g.Cells().Filter(func(c grid.Cell) bool {
		return g.Get(c) == "."
	})
	for openCells.Next() {
		cell := openCells.Value()
		for _, neighbour := range g.NeighboursAdjacent(cell, false) {
			if g.Get(neighbour) == "." {
				gra.AddCost(g.OffsetFromCell(cell), g.OffsetFromCell(neighbour), 1)
			}
		}
	}

	_, dist := graph.ShortestPath(gra,
		g.OffsetFromCell(grid.Cell{1, 1}),
		g.OffsetFromCell(target))
	*p1 = strconv.Itoa(int(dist))

	ttlGrid := grid.NewGrid[int](gridSize, gridSize)
	traverse(gra, ttlGrid.OffsetFromCell(grid.Cell{1, 1}), 51, ttlGrid)
	*p2 = strconv.Itoa(ttlGrid.FilterByValue(util.Neq(0)).Count())
}

func traverse(gra *graph.Mutable, node int, ttl int, g grid.Grid[int]) {
	if ttl == 0 {
		return
	}
	cell := g.CellFromOffset(node)
	cellValue := g.Get(cell)
	if ttl <= cellValue {
		// current path is no better than previous attempts, so give up
		return
	} else if ttl > cellValue {
		// current path is better, so overwrite previous value
		g.Set(cell, ttl)
		gra.Visit(node, func(w int, c int64) bool {
			traverse(gra, w, ttl-1, g)
			return false
		})
	}
}
