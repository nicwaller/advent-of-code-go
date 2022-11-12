package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"fmt"
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
	flashgrid := grid.TransformGrid[int](g, func(val int) bool {
		return false
	})

	incr := func(v int) int { return v + 1 }

	//For 100 steps
	firstSynchrony := -1
	flashesInFirst100Steps := 0
Forever:
	for step := 1; ; step++ {
		// First, the energy level of each octopus increases by 1.
		g.MapAll(incr)

		// Then, any octopus with an energy level greater than 9 flashesInFirst100Steps.
		// This increases the energy level of all adjacent octopuses by 1,
		// including octopuses that are diagonally adjacent.
		flashesInStep := 0
		flashgrid.Fill(false)
		var anyDidFlash bool
		for {
			anyDidFlash = false
			for cellIterator := g.Cells(); cellIterator.Next(); {
				cell := cellIterator.Value()
				v := g.Get(cell)
				if v > 9 {
					if flashgrid.Get(cell) {
						// it already flashed.
						continue
					}
					// the octopus flashesInFirst100Steps!!!
					if step <= 100 {
						flashesInFirst100Steps++
					}
					flashesInStep++
					flashgrid.Set(cell, true)
					anyDidFlash = true
					neighbours := g.NeighboursSurround(cell, true)
					g.MapIter(incr, iter.ListIterator[grid.Cell](&neighbours))
				}
			}
			if anyDidFlash {
				continue
			} else {
				if flashesInStep == g.RowCount()*g.ColumnCount() {
					fmt.Println("synchrony acheived")
					firstSynchrony = step
					break Forever
				}
				break
			}
		}
		// Finally, any octopus that flashed during this step has its energy level set to 0,
		// as it used all of its energy to flash.
		for tiredOcto := flashgrid.Filter(func(didFlash bool) bool { return didFlash }); tiredOcto.Next(); {
			g.Set(tiredOcto.Value(), 0)
		}
		//fmt.Printf("After step %d:\n", step)
		//g.Print()
	}
	//  - grid.items.filter(>9 and not flashed)
	//  - break if no result
	//  - flashNeighbours(diag=yes)
	// - grid.items.filter(>9).map(zero)
	assert.EqualAny(flashesInFirst100Steps, []int{1640}, "flashesInFirst100Steps")
	assert.EqualAny(firstSynchrony, []int{312}, "firstSynchrony")
	fmt.Printf("synchrony in step %d\n", firstSynchrony)
	return flashesInFirst100Steps
}

func part2(g grid.Grid[int]) int {
	//assert.EqualAny(basinMultiplyResult, []int{1134, 1023660}, "basinMultiplyResult")
	return -1
}
