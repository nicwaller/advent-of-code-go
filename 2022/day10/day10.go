package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
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

// decode a sequence of macro-instructions into microcode instructions
// that take exactly one cycle to process
func decode(ite iter.Iterator[string], next chan<- string) {
	go func() {
		for ite.Next() {
			line := ite.Value()
			switch strings.Fields(line)[0] {
			case "noop":
				next <- line
			case "addx":
				next <- "noop"
				next <- line
			default:
				panic(line)
			}
		}
		close(next)
	}()
}

func run(p1 *string, p2 *string) {
	crt := grid.NewGrid[string](6, 40)
	crt.Fill(".")

	X := 1
	significantSignals := make([]int, 6)
	microcode := make(chan string)
	decode(aoc.InputLinesIterator(), microcode)
	cycle := 0

runLoop:
	for {
		crtX := cycle % 40
		crtY := (cycle / 40) % 6
		if util.IntAbs(crtX-X) <= 1 {
			crt.Set(grid.Cell{crtY, crtX}, "#")
		}

		// This is an uncomfortable, weird place for the increment.
		// The increment is needed after CRT, but before sigSig calculation.
		cycle++

		if (cycle-20)%40 == 0 {
			significantSignals[(cycle-20)/40] = cycle * X
		}

		uInst, channelStillOpen := <-microcode
		if !channelStillOpen {
			break runLoop
		}

		switch strings.Fields(uInst)[0] {
		case "noop":
			// pass
		case "addx":
			X += util.UnsafeAtoi(strings.Fields(uInst)[1])
		default:
			panic(uInst)
		}
	}

	*p1 = strconv.Itoa(f8l.Sum(significantSignals))
	*p2 = aoc.Debannerize(crt, "#")
}
