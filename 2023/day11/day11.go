package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2023, 11)
	aoc.Test(run, "sample.txt", "374", "")
	aoc.Test(run, "input.txt", "9312968", "597714117556")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	g := expandUniverse(aoc.InputGridRunes())
	sum := 0
	galaxies := g.FilterByValue(isGalaxy).List()
	for g1 := 0; g1 < len(galaxies); g1++ {
		for g2 := g1 + 1; g2 < len(galaxies); g2++ {
			sum += grid.ManhattanDistance(galaxies[g1], galaxies[g2])
		}
	}

	g = aoc.InputGridRunes()
	xSM, ySM := expandUniverseScale(g, 1000000-1)
	sum2 := 0
	galaxies = g.FilterByValue(isGalaxy).List()
	for g1 := 0; g1 < len(galaxies); g1++ {
		for g2 := g1 + 1; g2 < len(galaxies); g2++ {
			z1 := util.Copy(galaxies[g1])
			z2 := util.Copy(galaxies[g2])
			z1[0] += ySM[z1[0]]
			z2[0] += ySM[z2[0]]
			z1[1] += xSM[z1[1]]
			z2[1] += xSM[z2[1]]
			sum2 += grid.ManhattanDistance(z1, z2)
		}
	}

	*p1 = strconv.Itoa(sum)
	*p2 = strconv.Itoa(sum2)
}

func expandUniverse(g grid.Grid[string]) grid.Grid[string] {
	xShift := 0
	yShift := 0
	xShiftMap := make([]int, g.ColumnCount())
	yShiftMap := make([]int, g.RowCount())
	for rowI := 0; rowI < g.RowCount(); rowI++ {
		if iter.ListIterator(g.Row(rowI)).Filter(isGalaxy).Count() == 0 {
			yShift++
		}
		yShiftMap[rowI] = yShift
	}
	for colI := 0; colI < g.ColumnCount(); colI++ {
		if iter.ListIterator(g.Column(colI)).Filter(isGalaxy).Count() == 0 {
			xShift++
		}
		xShiftMap[colI] = xShift
	}

	s := g.All()
	s[0].Terminus *= 2
	s[1].Terminus *= 2
	gBig := grid.NewGridFromSlice[string](s)
	gBig.Fill(".")
	g.FilterByValue(isGalaxy).Each(func(cell grid.Cell) {
		gBig.Set(grid.Cell{
			cell[0] + yShiftMap[cell[0]],
			cell[1] + xShiftMap[cell[1]],
		}, "#")
	})

	return gBig
}

func expandUniverseScale(g grid.Grid[string], scale int) (xShiftMap []int, yShiftMap []int) {
	xShift := 0
	yShift := 0
	xShiftMap = make([]int, g.ColumnCount())
	yShiftMap = make([]int, g.RowCount())
	for rowI := 0; rowI < g.RowCount(); rowI++ {
		if isEmptySpace(g.Row(rowI)) {
			yShift++
		}
		yShiftMap[rowI] = yShift * scale
	}
	for colI := 0; colI < g.ColumnCount(); colI++ {
		if isEmptySpace(g.Column(colI)) {
			xShift++
		}
		xShiftMap[colI] = xShift * scale
	}
	return xShiftMap, yShiftMap
}

func isGalaxy(s string) bool {
	return s == "#"
}

func isEmptySpace(scan []string) bool {
	return iter.ListIterator(scan).Filter(isGalaxy).Count() == 0
}
