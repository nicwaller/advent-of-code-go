package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
)

func main() {
	aoc.Day(17)
	aoc.Test(run, "sample.txt", "45", "")
	//aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	f := util.NumberFields(aoc.InputString())
	target := grid.SliceEnclosing(
		grid.Cell{f[0], f[2]},
		grid.Cell{f[1], f[3]},
	)

	fmt.Printf("target: %d < x < %d\n", f[0], f[1])

	var finalX = 0
	launchPossibilitiesX := make([]int, 0)
	for launchVX := 0; finalX <= f[1]; launchVX++ {
		finalX += launchVX
		if finalX > f[0] && finalX < f[1] {
			launchPossibilitiesX = append(launchPossibilitiesX, launchVX)
		}
	}
	fmt.Println(launchPossibilitiesX)

	maxVY := 0
	for _, vx := range launchPossibilitiesX {
		for vy := 0; vy < 100; vy++ {
			if simulate(vx, vy, target) {
				fmt.Printf("yes vx=%d, vy=%d\n", vx, vy)
				maxVY = util.IntMax(maxVY, vy)
			}
		}
	}
	peakY := iter.Range(0, maxVY+1).Reduce(util.IntSum, 0)
	fmt.Println(maxVY)
	fmt.Println(peakY)

	*p1 = strconv.Itoa(peakY)
	//*p2 = strconv.Itoa(p2)
}

func simulate(vx1 int, vy1 int, target grid.Slice) bool {
	x := 0
	y := 0
	vx := vx1
	vy := vy1

	yTrace := []int{y}
	for {
		x += vx
		y += vy
		yTrace = append(yTrace, y)
		if y < target[1].Origin {
			return false
		} else if target.Contains([]int{x, y}) {
			return true
		}
		if vx > 0 {
			vx--
		}
		vy--
	}
}
