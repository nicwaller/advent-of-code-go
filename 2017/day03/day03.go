package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"math"
	"strconv"
)

func main() {
	aoc.Select(2017, 3)
	aoc.TestLiteral(run, "12", "3", "")
	aoc.TestLiteral(run, "23", "2", "")
	aoc.TestLiteral(run, "1024", "31", "")
	aoc.TestLiteral(run, "1024", "31", "")
	//aoc.Test(run, "input.txt", "552", "330785")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	dirs := [][]int{
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 0},
	}

	target := util.UnsafeAtoi(aoc.InputString())
	targetCell := make([]int, 2)

	g := grid.NewGrid[int](1, 1)
	g2 := grid.NewGrid[int](1, 1)
	scale := int(math.Round(math.Sqrt(float64(target))))
	g.Grow(scale, 0)
	g2.Grow(scale, 0)
	dirV := 0
	steps := 1
	n := 1
	pos := grid.Cell{0, 0}
	g.Set(pos, n)
	g2.Set(pos, 1)
	var targetBigVal int
	for n < target {
		for j := 0; j < 2; j++ {
			for i := 1; i <= steps; i++ {
				// part1
				n++
				util.VecAdd(pos, dirs[dirV])
				g.Set(pos, n)
				if n == target {
					copy(targetCell, pos)
				}

				// part2
				neighbourCells := g2.NeighboursSurround(pos, false)
				neighbourVals := f8l.Map(neighbourCells, g2.Get)
				bigVal := f8l.Sum(neighbourVals)
				g2.Set(pos, bigVal)
				if targetBigVal == 0 && bigVal > target {
					targetBigVal = bigVal
				}
			}
			dirV = (dirV + 1) % len(dirs)
		}
		steps++
	}
	//fmt.Printf("Target %v found at %v\n", target, targetCell)

	*p1 = strconv.Itoa(grid.ManhattanDistance(grid.Cell{0, 0}, targetCell))
	*p2 = strconv.Itoa(targetBigVal)
}
