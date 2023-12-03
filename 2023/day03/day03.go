package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 3)
	aoc.Test(run, "sample.txt", "4361", "")
	aoc.Test(run, "input.txt", "533784", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	g := aoc.InputGridRunes()
	//g.Grow(1, ".") // TODO: grow is very confusing because indexing starts at -1
	g.Print()
	bounds := g.All() // TODO: maybe alias as grid.Bounds()
	// g.Contains(cell)?
	sum := 0
	//gs := grid.SliceEnclosing(
	//	grid.Cell{0, 0},
	//	grid.Cell{g.RowCount() - 1, g.ColumnCount() - 1},
	//)
	partNumbers := set.New[int]()
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
				fmt.Println(num)

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
						break
					}
				}

				if isPartNumber {
					sum += num
				}
			}
			// identify spans of numbers
			//if spanStart == nil {
			//	if isDigit(cv) {
			//		gotIt := x
			//		spanStart = &gotIt
			//		spanEnd = nil
			//	}
			//} else {
			//	if isDigit(cv) {
			//		// extending the number
			//		gotIt := x
			//		spanEnd = &gotIt
			//	} else {
			//		// number is terminated
			//		scanNow = true
			//	}
			//}
			//if spanStart != nil && spanEnd == nil && x+1 == g.ColumnCount() {
			//	scanNow = true
			//	gg := g.ColumnCount() - 1
			//	spanEnd = &gg
			//}
			//if scanNow {
			//	digits := make([]string, 0)
			//	for i := *spanStart; i <= *spanEnd; i++ {
			//		cc := grid.Cell{y, i}
			//		digits = append(digits, g.Get(cc))
			//	}
			//	digStr := strings.Join(digits, "")
			//	digVal := util.UnsafeAtoi(digStr)
			//	fmt.Println(digVal)
			//	s := grid.SliceEnclosing(
			//		grid.Cell{y - 1, *spanStart - 1},
			//		grid.Cell{y + 1, *spanEnd + 1},
			//	)
			//	s, _ = s.Intersect(gs)
			//	for sc := s.Cells(); sc.Next(); {
			//		scv := g.Get(sc.Value())
			//		if isSymbol(scv) {
			//			//fmt.Println(digVal)
			//			sum += digVal
			//			partNumbers.Insert(digVal)
			//			break
			//		}
			//	}
			//	spanStart = nil
			//	spanEnd = nil
			//	scanNow = false
			//}

			// for each span, check all neighbours

		}
	}
	sumAlt := f8l.Sum(partNumbers.Items())
	_ = sumAlt

	*p1 = strconv.Itoa(sum)
	*p2 = strconv.Itoa(0)
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
