package main

import (
	"fmt"
	"iter"
	"slices"
	"strconv"

	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid2"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2025, 7)
	aoc.Test(run, "sample.txt", "21", "")
	aoc.Test(run, "input.txt", "1516", "")
	aoc.Run(run)
}

const (
	EmptySpace       = '.'
	StartingManifold = 'S'
	TachyonBeam      = '|'
	BeamSplitter     = '^'
)

type Beam = iter.Seq[grid2.Cell[rune]]

func run(p1 *string, p2 *string) {
	g := grid2.NewGridFromString(aoc.InputString())
	*p1 = strconv.Itoa(approach1(g))
}

func approach1(g *grid2.Grid[rune]) int {
	beams := make([]Beam, 0)

	launchBeam := func(src grid2.Cell[rune]) {
		switch src.Get() {
		case EmptySpace:
			src.Set(TachyonBeam)
			beams = append(beams, src.Ray(grid2.Down))
		case TachyonBeam:
			// launch failed, does nothing
			return
		case BeamSplitter:
			panic("I did not expect this")
		default:
			panic("what other possibility is there?")
		}
	}

	// first, find the starting point "S" and launch the tachyon beam
	start := util.IterNext(g.Select().ValueEquals(StartingManifold))
	launchBeam(start.Neighbour(grid2.Down))

	// next, perpetuate the tachyon beams
	for keepGoing := true; keepGoing; {
		keepGoing = false
		for _, beam := range beams {
			nextCell := util.IterNext(beam)
			if nextCell == nil {
				// this beam is complete
				continue
			} else {
				// at least one beam is continuing
				keepGoing = true
			}
			ncv := nextCell.Get()
			switch ncv {
			case EmptySpace:
				// propagate the tachyon beam
				nextCell.Set(TachyonBeam)
			case BeamSplitter:
				// halt this beam
				_ = slices.Collect(beam)
				// make two more beams
				launchBeam(nextCell.Neighbour(grid2.Left))
				launchBeam(nextCell.Neighbour(grid2.Right))
			case TachyonBeam:
				// do nothing, already a beam here
				continue
			default:
				panic(fmt.Errorf("cell contains %q", ncv))
			}
		}
	}

	// NOTE: I'm not trying to find "how many beams split"
	// I'm trying to find "how many beam splitters are activated"
	// in other words, finding beam splitters that never have a beam reach them.
	// So let's find all the splitters that don't have a tachyon beam | directly above.

	splits := 0
	for splitterCell := range g.Select().ValueEquals(BeamSplitter) {
		aboveCell := splitterCell.Neighbour(grid2.Up)
		if aboveCell.Get() == TachyonBeam {
			splits++
		}
	}

	return splits
}
