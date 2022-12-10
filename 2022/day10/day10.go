package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/queue"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 10)
	aoc.Test(run, "sample.txt", "13140", "")
	aoc.Test(run, "input.txt", "14860", "RGZEHURK")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	crt := grid.NewGrid[string](6, 40)
	crt.Fill(".")

	X := 1
	significantSignals := make([]int, 6)
	lines := aoc.InputLinesIterator()
	microQ := queue.New[string](10)
	for cycle := 1; cycle <= 250; cycle++ {
		signalStrength := cycle * X
		if (cycle-20)%40 == 0 {
			significantSignals[(cycle-20)/40] = signalStrength
		}
		//fmt.Printf("cycle:%d, X:%d, str:%d\n",
		//	cycle, X, signalStrength)

		cycleOffOne := cycle - 1
		crtX := cycleOffOne % 40
		if util.IntAbs(crtX-X) <= 1 {
			crtY := (cycleOffOne / 40) % 6
			crt.Set(grid.Cell{crtY, crtX}, "#")
		}

		if microQ.Length() < 1 {
			lines.Next()
			nextInst := lines.Value()
			if nextInst == "noop" {
				_ = microQ.Push("noop")
			} else {
				_ = microQ.Push("noop")
				_ = microQ.Push(nextInst)
			}
		}

		uInst := util.Must(microQ.Pop())
		operator := strings.Fields(uInst)[0]

		switch operator {
		case "noop":
			// pass
		case "addx":
			operand := util.UnsafeAtoi(strings.Fields(uInst)[1])
			X += operand
		default:
			panic(operator)
		}
	}
	*p1 = strconv.Itoa(f8l.Sum(significantSignals))
	crt.Print()
	*p2 = aoc.Debannerize(crt, "#")
}
