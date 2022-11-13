package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"fmt"
	"strings"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

type fileType struct {
	dots  []grid.Cell
	folds []fold
}

type fold struct {
	axis     string
	position int
}

func parseFile() fileType {
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

	return fileType{
		dots:  dots,
		folds: folds,
	}
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC")
			// FIXME: why are the offsets non-zero?
			fmt.Printf("Current fold: %v\n", at)
			fmt.Printf(" origin:      %v\n destination: %v\n", foldOrigin, foldDestination)
			srcOffset := g.OffsetFromCell(foldOrigin)
			dstOffset := g.OffsetFromCell(foldDestination)
			fmt.Printf(" src:  %v\n dst:  %v\n", srcOffset, dstOffset)
		}
	}()

	s := g.All()
	s[dimension].Origin = at.position
	for cellIter := s.Cells(); cellIter.Next(); {
		foldOrigin = cellIter.Value()
		foldDestination = make([]int, len(foldOrigin))
		copy(foldDestination, foldOrigin)
		foldDestination[dimension] = at.position - (foldOrigin[dimension] - at.position)
		//fmt.Printf("%v -> %v\n", foldOrigin, foldDestination)
		if g.Get(foldOrigin) == "#" {
			g.Set(foldDestination, "#")
			g.Set(foldOrigin, ".")
		}
	}
	//g.FillSlice(".", s)
	//fmt.Println(at)
	//fmt.Println(s)
}

func part1(input fileType) int {
	fix := append(input.dots, grid.Cell{0, 0})
	g := grid.NewGridFromSlice[string](grid.SliceEnclosing(fix...))
	g.Fill(".")
	for _, dot := range input.dots {
		g.Set(dot, "#")
	}

	// First Fold
	foldInPlace(g, input.folds[0])
	visibleDots := g.Filter(func(v string) bool { return v == "#" }).Count()
	//assert.EqualAny(visibleDots, []int{759}, "visible dots")

	// Remaining Folds
	fmt.Println("folding at home...")
	for idx, f := range input.folds[1:] {
		fmt.Printf("starting fold #%d\n", idx)
		foldInPlace(g, f)
	}
	AoCPrint(g)
	return visibleDots
}

func part2(g fileType) int {
	//assert.EqualAny(basinMultiplyResult, []int{1134, 1023660}, "basinMultiplyResult")
	return -1
}

func AoCPrint(g grid.Grid[string]) {
	fmt.Println("")
	for rowIndex := 0; rowIndex < 50; rowIndex++ {
		fmt.Println(strings.Join(g.Row(rowIndex), ""))
	}
	fmt.Println("")
}
