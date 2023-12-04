package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/iterc"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 3)
	aoc.Test(run, "sample.txt", "4361", "467835")
	aoc.Test(run, "input.txt", "533784", "78826761")
	//aoc.Run(run)
	aoc.Out()
}

type Coord2D struct {
	y int
	x int
}

func run(p1 *string, p2 *string) {
	g := aoc.InputGridRunes()
	bounds := g.All()
	sum := 0
	maybeGears := make(map[Coord2D][]int)
	for y := 0; y < g.RowCount(); y++ {
		for x := 0; x < g.ColumnCount(); x++ {
			c := grid.Cell{y, x}
			cv := g.Get(c)

			x0 := x
			x1 := -1
			if isDigit(cv) {
				for {
					if x == g.ColumnCount() {
						x1 = x - 1
						break
					}
					c := grid.Cell{y, x}
					cv := g.Get(c)
					if isDigit(cv) {
						x1 = x
					} else {
						break
					}
					x++
				}
			}
			if x0 >= 0 && x1 >= 0 {
				// we got a full number
				s := grid.SliceEnclosing(grid.Cell{y, x0}, grid.Cell{y, x1})
				digitIter := iter.Transform[grid.Cell, string](s.Cells(), g.Get)
				digits := digitIter.List()
				numStr := strings.Join(digits, "")
				num := util.UnsafeAtoi(numStr)

				// but is it a part number?
				isPartNumber := false
				s0 := grid.SliceEnclosing(grid.Cell{y - 1, x0 - 1}, grid.Cell{y + 1, x1 + 1})
				for nci := s0.Cells(); nci.Next(); {
					nc := nci.Value()
					if !bounds.Contains(nc) {
						continue
					}
					ncv := g.Get(nc)
					if isSymbol(ncv) {
						isPartNumber = true
						if ncv == "*" {
							// it's a gear!
							maybeGears[Coord2D{
								y: nc[0],
								x: nc[1],
							}] = append(maybeGears[Coord2D{
								y: nc[0],
								x: nc[1],
							}], num)
						}
					}
				}

				if isPartNumber {
					sum += num
				}
			}
		}
	}

	gearSum := 0
	iterc.MapIterator(maybeGears).Filter(func(k iterc.KV[Coord2D, []int]) bool {
		return len(k.Value) == 2
	}).ForEach(func(k iterc.KV[Coord2D, []int]) {
		ratio := k.Value[0] * k.Value[1]
		gearSum += ratio
	})

	*p1 = strconv.Itoa(sum)
	*p2 = strconv.Itoa(gearSum)
}

func isSymbol(s string) bool {
	if s == "." {
		return false
	}
	if isDigit(s) {
		return false
	}
	return true
}

func isDigit(s string) bool {
	return strings.ContainsAny(s, "0123456789")
}
