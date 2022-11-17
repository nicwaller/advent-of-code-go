package main

import (
	"advent-of-code/lib/aoc"
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
		*p1 = strconv.Itoa(exec(program))
		return
	} else {
		*p1 = strconv.Itoa(execArgs(program, 12, 2))
	}

Permutations:
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			prog := make([]int, len(program))
			copy(prog, program)
			result := execArgs(program, noun, verb)
			if result == 19690720 {
				*p2 = strconv.Itoa(100*noun + verb)
				break Permutations
			}
		}
	}
}

func execArgs(progOrig []int, arg1 int, arg2 int) int {
	prog := make([]int, len(progOrig))
	copy(prog, progOrig)
	prog[1] = arg1
	prog[2] = arg2
	return exec(prog)
}

func exec(progOrig []int) int {
	prog := make([]int, len(progOrig))
	copy(prog, progOrig)

	ptr := 0
	for {
		opcode := prog[ptr]
		if opcode == 99 {
			break
		}
		op1 := prog[prog[ptr+1]]
		op2 := prog[prog[ptr+2]]
		register := prog[ptr+3]
		switch opcode {
		case 1: // add
			prog[register] = op1 + op2
		case 2: // mult
			prog[register] = op1 * op2
		default:
			panic(opcode)
		}
		ptr += 4
	}
	return prog[0]
}
