package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"math"
	"strconv"
)

func main() {
	aoc.Select(2019, 8)
	aoc.Test(run, "sample.txt", "", "")
	aoc.Test(run, "sample2.txt", "", "")
	aoc.Test(run, "input.txt", "1088", "LGYHB")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	var rows int
	var cols int
	input := aoc.InputString()
	switch len(input) {
	case 15000:
		rows = 6
		cols = 25
	case 16:
		rows = 2
		cols = 2
	case 12:
		rows = 2
		cols = 3
	default:
		panic(len(input))
	}
	layersStr := iter.StringIterator(input, rows*cols).List()
	minn := math.MaxInt32
	score := 0
	g := grid.NewGrid[string](rows, cols)
	g.Fill("2")
	g.Values()

	for _, layer := range layersStr {
		a := analyze.CountDistinct(iter.StringIterator(layer, 1).List())
		if a["0"] < minn {
			minn = a["0"]
			score = a["1"] * a["2"]
		}
		for i, curVal := range g.Values() {
			if curVal == "2" {
				if nextVal := layer[i : i+1]; nextVal != "2" {
					g.Values()[i] = nextVal
				}
			}
		}
	}
	*p1 = strconv.Itoa(score)

	if g.RowCount() == 6 {
		g.Replace("0", ".")
		g.Replace("1", "█")
		//g.Print()
		*p2 = aoc.Debannerize(g, "█")
	}
}
