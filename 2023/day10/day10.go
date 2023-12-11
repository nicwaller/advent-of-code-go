package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"advent-of-code/lib/stack"
	"strconv"
)

func main() {
	aoc.Select(2023, 10)
	//aoc.Test(run, "sample.txt", "4", "")
	//aoc.Test(run, "sample2.txt", "8", "")
	//aoc.Test(run, "sample3.txt", "", "4")
	//aoc.Test(run, "sample4.txt", "", "8")
	//aoc.Test(run, "sample5.txt", "", "10")
	//aoc.Test(run, "input.txt", "6640", "")
	// 701 too high
	// 692 is too high
	// 661 is too high
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	// find the tile in the loop that is farthest from the starting position
	// find the tile that would take the longest number of steps along the loop to reach from the starting point
	// regardless of which way around the loop the animal went

	g := aoc.InputGridRunes()
	g.Grow(1, ".")
	S := g.Filter(func(cell grid.Cell, s string) bool {
		return s == "S"
	}).List()[0]
	//fmt.Println(S)

	g2 := grid.NewGridFromSlice[int](g.All())
	g3 := grid.NewGridFromSlice[string](g.All())
	g3.Fill("I")
	g3.Set(S, "S")

	//fwd := iter.Chain(iter.ListIterator([]grid.Cell{S}), pipeForward(g, S)).List()
	fwd := pipeForward(g, S).List()
	fwdIter := iter.ListIterator(fwd)
	for i := 1; fwdIter.Next(); i++ {
		cell := fwdIter.Value()
		g2.Set(cell, i)
		g3.Set(cell, g.Get(cell))
	}
	bkwdIter := pipeBackward(fwd)
	for i := 1; bkwdIter.Next(); i++ {
		cell := bkwdIter.Value()
		//fmt.Println("+", cell)
		if g2.Get(cell) > i {
			g2.Set(cell, i)
		}
		//g3.Set(cell, ".")
		g3.Set(cell, g.Get(cell))
	}

	highest := 0
	g2.ValuesIterator().Each(func(i int) {
		highest = max(highest, i)
	})

	// this is not quite correct either
	//g3.FloodFill(grid.Cell{4, 8}, func(v string) bool {
	//	return v == " "
	//}, "i")
	g3.FloodFill(grid.Cell{0, 0}, func(v string) bool {
		return v == "I"
	}, "O")
	g3.Print()

	enclosed := g3.FilterByValue(func(s string) bool {
		return s == "I"
	}).Count()

	*p1 = strconv.Itoa(highest)
	*p2 = strconv.Itoa(enclosed)
}

func cellsEqual(a, b grid.Cell) bool {
	// TODO: this is a stupid way to compare cell equality
	return grid.CellHash(a) == grid.CellHash(b)
}

func contains(haystack []string, needle string) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func pipeForward(g grid.Grid[string], origin grid.Cell) iter.Iterator[grid.Cell] {
	cur := origin

	//ups := set.New("|", "7", "F")
	//downs := set.New("|", "L", "J") // TODO
	//lefts := set.New("-", "L", "F")
	//rights := set.New("-", "7", "J")

	seen := set.New[int]()
	return iter.Iterator[grid.Cell]{
		Next: func() bool {
			nb := g.NeighboursAdjacent(cur, false)
			up := nb[0]
			down := nb[1]
			left := nb[2]
			right := nb[3]
			upV := g.Get(up)
			downV := g.Get(down)
			leftV := g.Get(left)
			rightV := g.Get(right)

			cv := g.Get(cur)

			if !seen.Contains(grid.CellHash(up)) {
				if contains([]string{"|", "L", "J", "S"}, cv) {
					if contains([]string{"|", "7", "F"}, upV) {
						cur = up
						goto found
					}
				}
			}

			if !seen.Contains(grid.CellHash(right)) {
				if contains([]string{"-", "L", "F", "S"}, cv) {
					if contains([]string{"-", "J", "7"}, rightV) {
						cur = right
						goto found
					}
				}
			}

			if !seen.Contains(grid.CellHash(down)) {
				if contains([]string{"|", "7", "F", "S"}, cv) {
					if contains([]string{"|", "L", "J"}, downV) {
						cur = down
						goto found
					}
				}
			}

			if !seen.Contains(grid.CellHash(left)) {
				if contains([]string{"-", "7", "J", "S"}, cv) {
					if contains([]string{"-", "L", "F"}, leftV) {
						cur = left
						goto found
					}
				}
			}

			return false

		found:
			seen.Insert(grid.CellHash(cur))
			return !cellsEqual(origin, cur)

			//if seen.Contains(grid.CellHash(cur)) {
			//	//panic(seen)
			//	fmt.Println("warning")
			//	time.Sleep(time.Second)
			//} else {
			//	seen.Insert(grid.CellHash(cur))
			//}

			//p := cur
			//cv := g.Get(cur)
			//
			//// TODO: am I sure this won't loop back upon itself? should I have set of seen?
			//if upVal := g.Get(up); ups.Contains(upVal) && !cellsEqual(prev, up) {
			//	cur = up
			//} else if rightVal := g.Get(right); rights.Contains(rightVal) && !cellsEqual(prev, right) {
			//	cur = right
			//} else if downVal := g.Get(down); downs.Contains(downVal) && !cellsEqual(prev, down) {
			//	cur = down
			//} else if leftVal := g.Get(left); lefts.Contains(leftVal) && !cellsEqual(prev, left) {
			//	cur = left
			//} else {
			//	return false
			//	//panic("Wtf")
			//}
			//
			//prev = p
			//
			////fmt.Println(origin, cur)

		},
		Value: func() grid.Cell {
			//fmt.Println(cur)
			//time.Sleep(200 * time.Millisecond)
			return cur
		},
	}
}

func pipeBackward(fwd []grid.Cell) iter.Iterator[grid.Cell] {
	sk := stack.FromSlice[grid.Cell](fwd)
	return sk.Iterator()
}
