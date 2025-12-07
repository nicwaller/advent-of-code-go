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
	aoc.Test(run, "sample.txt", "4277556", "3263827")
	aoc.Test(run, "input.txt", "6417439773370", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	h := parse()
	*p1 = strconv.Itoa(solve(h, false))
	*p2 = strconv.Itoa(solve(h, true))
}

type Homework = []Problem

func parse() Homework {
	var hw Homework

	// character alignment is of critical importance in part two
	// and not all lines are the same length
	// but grid.FromString() handles all that correctly.
	inputGrid := grid.FromString(aoc.InputString())

	// We use operator positions in the last row to select blocks of numbers.
	operatorsRow := inputGrid.Row(inputGrid.RowCount() - 1)

	// Operators are left-aligned within their block.
	// So scanning right-to-left means every time we see an operator,
	// we're done scanning a whole block.
	x2 := len(operatorsRow) - 1
	yTop := 0
	yBottom := inputGrid.RowCount() - 2
	for x1 := len(operatorsRow) - 1; x1 >= 0; x1-- {
		op := operatorsRow[x1]
		switch op {
		case "", " ":
			// keep scanning until we find an operator
			// it's normal and expected for some columns to not have an operator
			continue

		case "+", "*":
			// we found an operator, so let's figure out the bounding box
			// that contains all the values of interest
			subSlice := grid.SliceEnclosing(
				grid.Cell{yTop, x1},
				grid.Cell{yBottom, x2},
			)
			// with that bounding box, we can extract that into a separate grid
			// this is helpful later on when we want to transpose the values
			g2str := inputGrid.SubGrid(subSlice)
			hw = append(hw, Problem{
				values:   g2str,
				operator: op,
			})
			// reset the watermark for the right edge boundary
			// also skip the empty column (optional, but nice)
			x2 = x1 - 2
		default:
			panic(fmt.Errorf("unknown operator %q", op))
		}
	}
	return hw
}

type Problem struct {
	values   grid.Grid[string]
	operator string
}

func (p Problem) Values(transpose bool) []int {
	g := p.values
	if transpose {
		g = g.Transpose()
	}

	values := make([]int, 0, g.RowCount())
	for rowIndex := 0; rowIndex < g.RowCount(); rowIndex++ {
		row := g.Row(rowIndex)
		rowStr := strings.TrimSpace(strings.Join(row, ""))
		if rowStr == "" {
			continue
		}
		v := util.UnsafeAtoi(rowStr)
		values = append(values, v)
	}

	return values
}

func solve(homework Homework, transpose bool) int {
	sum := 0

	for _, problem := range homework {
		intValues := problem.Values(transpose)
		switch problem.operator {
		case "+":
			result := f8l.Reduce(intValues, 0, func(a int, b int) int {
				return a + b
			})
			sum += result
		case "*":
			result := f8l.Reduce(intValues, 1, func(a int, b int) int {
				return a * b
			})
			sum += result
		default:
			panic(fmt.Errorf("unknown operator %q", problem.operator))
		}

	}

	return sum
}
