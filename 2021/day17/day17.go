package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
)

func main() {
	aoc.Select(2021, 17)
	aoc.Test(run, "sample.txt", "45", "112")
	aoc.Test(run, "input.txt", "4851", "1739")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	f := util.NumberFields(aoc.InputString())
	targetX0 := f[0]
	targetX1 := f[1]
	targetY0 := f[2]
	targetY1 := f[3]

	count := 0
	highestApex := 0

	for vx1 := 0; vx1 < 300; vx1++ {
		for vy1 := -100; vy1 < 100; vy1++ {
			vx := vx1
			vy := vy1
			x := 0
			y := 0
			apex := 0
			hitTarget := false
		Simulation:
			for {
				if vy == 0 {
					apex = y
				}
				x += vx
				y += vy
				if targetX0 <= x && x <= targetX1 {
					if targetY0 <= y && y <= targetY1 {
						hitTarget = true
						count++
						break Simulation
					}
				}
				if x > targetX1 {
					break Simulation
				}
				if y < util.IntMin(targetY0, targetY1) {
					break Simulation
				}
				if vx > 0 {
					vx--
				}
				vy--
			}
			if hitTarget {
				fmt.Printf("Solution: vx=%d, vy=%d\n", vx1, vy1)
				highestApex = util.IntMax(highestApex, apex)
			}
		}
	}
	//peakY := iter.Range(0, maxVY+1).Reduce(util.IntSum, 0)
	//fmt.Println(maxVY)
	//fmt.Println(peakY)

	*p1 = strconv.Itoa(highestApex)
	*p2 = strconv.Itoa(count)
}
