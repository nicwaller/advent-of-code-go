package aoc

import (
	"advent-of-code/lib/grid"
	"fmt"
	"strconv"
	"strings"
)

func bannerAlphabet() map[string]grid.Grid[bool] {
	letters := map[string]string{}
	letters["B"] = `
###..
#..#.
###..
#..#.
#..#.
###..
`
	letters["C"] = `
.##..
#..#.
#....
#....
#..#.
.##..
`
	letters["E"] = `
####.
#....
###..
#....
#....
####.
`
	letters["F"] = `
####.
#....
###..
#....
#....
#....
`
	letters["G"] = `
.##..
#..#.
#....
#.##.
#..#.
.###.
`
	letters["H"] = `
#..#.
#..#.
####.
#..#.
#..#.
#..#.
`
	letters["K"] = `
#..#.
#.#..
##...
#.#..
#.#..
#..#.
`
	letters["L"] = `
#....
#....
#....
#....
#....
####.
`
	letters["R"] = `
###..
#..#.
#..#.
###..
#.#..
#..#.
`
	letters["U"] = `
#..#.
#..#.
#..#.
#..#.
#..#.
.##..
`
	letters["Y"] = `
#...#
#...#
.#.#.
..#..
..#..
..#..
`
	letters["Z"] = `
####.
...#.
..#..
.#...
#....
####.
`

	letterGrids := make(map[string]grid.Grid[bool])
	for k, v := range letters {
		v = strings.ReplaceAll(strings.TrimSpace(v), "\n", "")
		if len(v) != 30 {
			fmt.Println(v[0:5], v[5:10], v[10:15], v[15:20], v[20:25], v[25:30])
			fmt.Println("-----")
			fmt.Println(v[30:])
			fmt.Printf("Expected 30 characters but got %d\n", len(v))
			panic(k)
		}
		g := grid.NewGrid[bool](6, 5)
		for i, _ := range v {
			if v[i:i+1] == "#" {
				g.Set(g.CellFromOffset(i), true)
			}
		}
		letterGrids[k] = g
	}
	return letterGrids
}

func Bannerize(s string, r string) grid.Grid[string] {
	bannerGrid := grid.NewGrid[string](6, len(s)*5)
	alpha := bannerAlphabet()
	text := strings.ToUpper(s)
	lookup := map[bool]string{
		false: " ",
		true:  r,
	}
	for i, _ := range text {
		c := text[i : i+1]
		letterGrid := alpha[c]
		for cellIter := letterGrid.Cells(); cellIter.Next(); {
			srcCell := cellIter.Value()
			dstCell := grid.Cell{srcCell[0], srcCell[1] + 5*i}
			bannerGrid.Set(dstCell, lookup[letterGrid.Get(srcCell)])
		}
	}
	return bannerGrid
}

func Debannerize(originalGrid grid.Grid[string], r string) string {
	var sb strings.Builder
	bannerGrid := grid.TransformGrid[string, bool](originalGrid, func(s string) bool { return s == r })
	if bannerGrid.RowCount() != 6 {
		panic("Banner grid must be height 6; got " + strconv.Itoa(bannerGrid.RowCount()))
	}
	if bannerGrid.ColumnCount()%5 != 0 {
		panic("Banner grid width must be a multiple of 5")
	}
	letterCount := bannerGrid.ColumnCount() / 5
	alphabet := bannerAlphabet()
	for l := 0; l < letterCount; l++ {
		// compare with each alphabet grid, one by one
	AlphabetScan:
		for letter, letterGrid := range alphabet {
			isMatch := true
			for cellIter := letterGrid.Cells(); cellIter.Next(); {
				letterCell := cellIter.Value()
				bannerCell := grid.Cell{letterCell[0], letterCell[1] + 5*l}
				if letterGrid.Get(letterCell) != bannerGrid.Get(bannerCell) {
					isMatch = false
					break
				}
			}
			if isMatch {
				sb.WriteString(letter)
				break AlphabetScan
			}
			// TODO: what if there is no match?
		}
	}
	return sb.String()
}
