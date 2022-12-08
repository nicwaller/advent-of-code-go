package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2022, 8)
	aoc.Test(run, "sample.txt", "21", "8")
	aoc.Test(run, "input.txt", "1816", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	g := aoc.InputGridNumbers()
	vg := g.Copy()
	vg.Fill(0)
	g.Print()
	cIter := g.Cells()
	visible := 0
	g.All()

	scenicScore := func(viewCell grid.Cell) int {
		iTreeHeight := g.Get(viewCell)
		tY := viewCell[0]
		tX := viewCell[1]

		lines := []iter.Iterator[grid.Cell]{
			grid.Line(viewCell, grid.Cell{tY, 0}),
			grid.Line(viewCell, grid.Cell{tY, g.Width() - 1}),
			grid.Line(viewCell, grid.Cell{0, tX}),
			grid.Line(viewCell, grid.Cell{g.Height() - 1, tX}),
		}
		vDists := make([]int, 4)
		for lineIndex, l := range lines {
			vals := f8l.Map(l.List(), g.Get)[1:]
			for _, v := range vals {
				if v < iTreeHeight {
					vDists[lineIndex]++
				} else {
					vDists[lineIndex]++
					break
				}
			}
		}
		score := f8l.Reduce(vDists, 1, util.IntProduct)
		return score
	}

	scenics := analyze.Box[int]{}
treeLoop:
	for cIter.Next() {
		treeCell := cIter.Value()
		treeHeight := g.Get(treeCell)
		tY := treeCell[0]
		tX := treeCell[1]

		scenics.Put(scenicScore(treeCell))

		lines := []iter.Iterator[grid.Cell]{
			grid.Line(treeCell, grid.Cell{tY, 0}),
			grid.Line(treeCell, grid.Cell{tY, g.Width() - 1}),
			grid.Line(treeCell, grid.Cell{0, tX}),
			grid.Line(treeCell, grid.Cell{g.Height() - 1, tX}),
		}
		for _, l := range lines {
			vals := iter.Map[grid.Cell, int](l, g.Get).List()[1:]
			if len(vals) == 0 {
				vg.Set(treeCell, 1)
				visible++
				continue treeLoop
			}
			lMax := analyze.Analyze(vals).Max
			if treeHeight > lMax {
				vg.Set(treeCell, 1)
				visible++
				continue treeLoop
			}
		}

		//var viz bool
		//
		//viz = true
		//for x := tX - 1; x >= 0; x-- {
		//	if g.Get(grid.Cell{tY, x}) >= treeHeight {
		//		// not visible this way
		//		viz = false
		//		break
		//	}
		//}
		//if viz {
		//	vg.Set(treeCell, 1)
		//	visible++
		//	continue
		//}
		//
		//viz = true
		//for x := tX + 1; x < g.Width(); x++ {
		//	if g.Get(grid.Cell{tY, x}) >= treeHeight {
		//		// not visible this way
		//		viz = false
		//		break
		//	}
		//}
		//if viz {
		//	vg.Set(treeCell, 1)
		//
		//	visible++
		//	continue
		//}
		//
		//viz = true
		//for y := tY - 1; y >= 0; y-- {
		//	if g.Get(grid.Cell{y, tX}) >= treeHeight {
		//		// not visible this way
		//		viz = false
		//		break
		//	}
		//}
		//if viz {
		//	vg.Set(treeCell, 1)
		//
		//	visible++
		//	continue
		//}
		//
		//viz = true
		//for y := tY + 1; y < g.Height(); y++ {
		//	if g.Get(grid.Cell{y, tY}) >= treeHeight {
		//		// not visible this way
		//		viz = false
		//		break
		//	}
		//}
		//if viz {
		//	vg.Set(treeCell, 1)
		//
		//	visible++
		//	continue
		//}

		continue treeLoop
	}
	vg.Print()
	*p1 = strconv.Itoa(visible)
	*p2 = strconv.Itoa(scenics.Max)
}
