package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/aoc/intcode"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2019, 2)
	aoc.Test(run, "sample.txt", "3500", "")
	aoc.Test(run, "input.txt", "5534943", "7603")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	f := strings.Split(aoc.InputString(), ",")
	program := f8l.Map(f, util.UnsafeAtoi)

	if aoc.InputString() == "1,9,10,3,2,3,11,0,99,30,40,50" {
		*p1 = strconv.Itoa(intcode.Exec(program))
		return
	} else {
		*p1 = strconv.Itoa(intcode.ExecArgs(program, []int{12, 2}))
	}

Permutations:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			prog := make([]int, len(program))
			copy(prog, program)
			result := intcode.ExecArgs(prog, []int{noun, verb})
			if result == 19690720 {
				*p2 = strconv.Itoa(100*noun + verb)
				break Permutations
			}
		}
	}
}
