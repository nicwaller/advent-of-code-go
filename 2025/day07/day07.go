package main

import (
	"strconv"

	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid2"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2025, 7)
	aoc.Test(run, "sample.txt", "21", "40")
	aoc.Test(run, "input.txt", "1516", "")
	aoc.Run(run)
}

const (
	EmptySpace    = '.'
	StartingPoint = 'S'
	TachyonBeam   = '|'
	BeamSplitter  = '^'
)

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(countActivatedBeamSplitters())
	*p2 = strconv.Itoa(countMultiverses())
}

func countActivatedBeamSplitters() int {
	g := grid2.NewGridFromString(aoc.InputString())

	// find the starting point "S"
	start := util.IterNext(g.Select().ValueEquals(StartingPoint))

	// then launch a single tachyon beam and keep track of how many beam splitters we encounter
	splitCount := 0
	propagateTachyonBeam(g, start.Neighbour(grid2.Down), &splitCount)

	return splitCount
}

// depth-first traversal is fun and easy!
func propagateTachyonBeam(g *grid2.Grid[rune], at grid2.Cell[rune], splitCount *int) {
	if !at.Coordinates().In(g.All().Bounds()) {
		return
	}
	switch at.Get() {
	case EmptySpace:
		at.Set(TachyonBeam)
		propagateTachyonBeam(g, at.Neighbour(grid2.Down), splitCount)
	case BeamSplitter:
		*splitCount++
		propagateTachyonBeam(g, at.Neighbour(grid2.Left), splitCount)
		propagateTachyonBeam(g, at.Neighbour(grid2.Right), splitCount)
	case TachyonBeam:
		// do nothing
		return
	case StartingPoint:
		panic("this should be impossible")
	}
}

func countMultiverses() int {
	g := grid2.NewGridFromString(aoc.InputString())
	particlesPerColumn := make([]int, g.All().Bounds().Dx()) // assumes origin at 0,0

	// first, find the starting point "S" and launch a single tachyon beam
	g.Select().ValueEquals(StartingPoint)(func(start grid2.Cell[rune]) bool {
		particlesPerColumn[start.Coordinates().X]++
		return false // halt search after finding the first manifold (there should only be one)
	})

	// for each beam splitter, update the count of multiverses
	g.Select().ValueEquals(BeamSplitter)(func(start grid2.Cell[rune]) bool {
		x := start.Coordinates().X
		particlesPerColumn[x-1] += particlesPerColumn[x]
		particlesPerColumn[x+1] += particlesPerColumn[x]
		particlesPerColumn[x] = 0
		return true
	})

	return f8l.Reduce(particlesPerColumn, 0, util.IntSum)
}
