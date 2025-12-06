package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2025, 6)
	aoc.Test(run, "sample.txt", "4277556", "6417439773370")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLines()
	rows := len(lines)
	columns := len(strings.Fields(lines[0]))

	g := grid.NewGrid[string](rows, columns)
	for y, row := range lines {
		for x, field := range strings.Fields(row) {
			g.Set(grid.Cell{y, x}, field)
		}
	}
	partOneSum := 0
	for x := 0; x < columns; x++ {
		column := g.Column(x)
		strValues := column[0 : len(column)-1]
		intValues := f8l.Map(strValues, util.UnsafeAtoi)
		operator := column[len(column)-1]
		switch operator {
		case "+":
			result := f8l.Reduce(intValues, 0, func(a int, b int) int {
				return a + b
			})
			partOneSum += result
		case "*":
			result := f8l.Reduce(intValues, 1, func(a int, b int) int {
				return a * b
			})
			partOneSum += result
		default:
			panic(fmt.Errorf("unknown operator %q", operator))
		}
	}
	*p1 = strconv.Itoa(partOneSum)
}
