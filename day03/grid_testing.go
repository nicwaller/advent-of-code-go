package main

import (
	"advent-of-code/day03/grid"
	"fmt"
)

func main() {
	//g := grid.NewGrid[rune](3, 5)
	g := grid.FromStringAsInt("01010\n10101")
	//g.Fill(0)
	//g.FillFunc2D(func(v rune, x int, y int) rune {
	//	return rune(x + y)
	//})
	g.Map(func(v int) int {
		return 3 * v
	})
	//g.TransposeInPlace()
	//reps := g.Replace(1, 3)
	//fmt.Printf("replaced %d occurrences\n", reps)
	fmt.Println(g)
	fmt.Printf("has %d rows of size %d\n", g.RowCount(), g.RowSize())
}
