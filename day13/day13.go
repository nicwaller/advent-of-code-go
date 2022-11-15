package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aoc.Day(13)
	aoc.Test(run, "sample.txt", "17", "")
	aoc.Test(run, "input.txt", "759", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLinesIterator()

	dots := make([]grid.Cell, 0)
	folds := make([]fold, 0)

	isNonEmptyString := func(line string) bool { return line != "" }
	for line := lines.TakeWhile(isNonEmptyString); line.Next(); {
		f := util.NumberFields(line.Value())
		dots = append(dots, grid.Cell{f[1], f[0]})
	}

	for lines.Next() {
		line := lines.Value()
		line = strings.Replace(line, "fold along ", "", 1)
		f := strings.Split(line, "=")
		folds = append(folds, fold{
			axis:     f[0],
			position: util.UnsafeAtoi(f[1]),
		})
	}

	fix := append(dots, grid.Cell{0, 0})
	g := grid.NewGridFromSlice[string](grid.SliceEnclosing(fix...))
	g.Fill(".")
	for _, dot := range dots {
		g.Set(dot, "#")
	}

	// First Fold
	foldInPlace(g, folds[0])
	visibleDots := g.Filter(func(v string) bool { return v == "#" }).Count()
	*p1 = strconv.Itoa(visibleDots)

	// Remaining Folds
	for _, f := range folds[1:] {
		foldInPlace(g, f)
	}

	// it wouldn't be easy to convert this into an actual string result
	AoCPrint(g)
}

type fold struct {
	axis     string
	position int
}

func foldInPlace(g grid.Grid[string], at fold) {
	dimension := -1
	switch at.axis {
	case "x":
		dimension = 1
	case "y":
		dimension = 0
	}
	_ = dimension

	var foldOrigin grid.Cell
	var foldDestination grid.Cell

	s := g.All()
	s[dimension].Origin = at.position
	for cellIter := s.Cells(); cellIter.Next(); {
		foldOrigin = cellIter.Value()
		foldDestination = make([]int, len(foldOrigin))
		copy(foldDestination, foldOrigin)
		foldDestination[dimension] = at.position - (foldOrigin[dimension] - at.position)
		if g.Get(foldOrigin) == "#" {
			g.Set(foldDestination, "#")
			g.Set(foldOrigin, ".")
		}
	}
}

func AoCPrint(g grid.Grid[string]) {
	fmt.Println("")
	for rowIndex := 0; rowIndex < 6; rowIndex++ {
		fmt.Println(strings.Join(g.Row(rowIndex)[:60], ""))
	}
	fmt.Println("")
}
