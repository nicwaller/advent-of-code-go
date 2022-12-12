package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2016, 12)
	//aoc.Test(run, "input.txt", "318009", "9227663")
	aoc.Run(run)
}

func runVM(registers map[string]int) {
	tape := aoc.InputLines()
	for pc := 0; pc < len(tape); pc++ {
		f := strings.Fields(tape[pc])
		op := f[0]
		operand := f[1]
		var operandV int
		if v, err := strconv.Atoi(operand); err == nil {
			operandV = v
		} else {
			operandV = registers[operand]
		}
		switch op {
		case "cpy":
			registers[f[2]] = operandV
		case "inc":
			registers[f[1]]++
		case "dec":
			registers[f[1]]--
		case "jnz":
			if operandV != 0 {
				pc += util.UnsafeAtoi(f[2])
				pc--
			}
		default:
			panic(op)
		}
	}
}

func run(p1 *string, p2 *string) {
	registers := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	}
	runVM(registers)
	*p1 = strconv.Itoa(registers["a"])

	registers = map[string]int{
		"a": 0,
		"b": 0,
		"c": 1,
		"d": 0,
	}
	runVM(registers)
	*p2 = strconv.Itoa(registers["a"])
}
