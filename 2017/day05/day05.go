package main

import (
	"advent-of-code/lib/aoc"
	"strconv"
)

func main() {
	aoc.Select(2017, 5)
	aoc.Test(run, "sample.txt", "5", "")
	//aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	tape := aoc.InputLinesInt()
	pc := 0
	steps := 0
	for {
		jmpOffset := tape[pc]
		tape[pc]++
		pc += jmpOffset
		steps++
		if pc < 0 || pc >= len(tape) {
			break
		}
	}
	*p1 = strconv.Itoa(steps)

	tape = aoc.InputLinesInt()
	pc = 0
	steps = 0
	for {
		jmpOffset := tape[pc]
		if jmpOffset >= 3 {
			tape[pc]--
		} else {
			tape[pc]++
		}
		pc += jmpOffset
		steps++
		if pc < 0 || pc >= len(tape) {
			break
		}
	}
	*p2 = strconv.Itoa(steps)
}
