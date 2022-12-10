package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"fmt"
	"golang.org/x/exp/maps"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2017, 8)
	aoc.Test(run, "sample.txt", "1", "10")
	//aoc.Test(run, "input.txt", "5752", "6366")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	a2 := analyze.Box[int]{}
	registers := make(map[string]int)
	for _, line := range aoc.InputLines() {
		ff := strings.Fields(line)
		r := ff[0]
		_, exist := registers[r]
		if !exist {
			registers[r] = 0
		}
		op := ff[1]
		operand := util.UnsafeAtoi(ff[2])
		condR := ff[4]
		condRVal, exist := registers[condR]
		if !exist {
			registers[condR] = 0
		}
		condOp := ff[5]
		condOperand := util.UnsafeAtoi(ff[6])
		switch condOp {
		case ">":
			if !(condRVal > condOperand) {
				continue
			}
		case "<":
			if !(condRVal < condOperand) {
				continue
			}
		case ">=":
			if !(condRVal >= condOperand) {
				continue
			}
		case "<=":
			if !(condRVal <= condOperand) {
				continue
			}
		case "==":
			if !(condRVal == condOperand) {
				continue
			}
		case "!=":
			if !(condRVal != condOperand) {
				continue
			}
		default:
			panic(condOp)
		}
		switch op {
		case "inc":
			registers[r] += operand
			a2.Put(registers[r])
		case "dec":
			registers[r] -= operand
			a2.Put(registers[r])
		default:
			panic(op)
		}
		fmt.Printf("%s = %d\n", r, registers[r])

		//fmt.Println(registers)
	}

	a := analyze.Analyze(maps.Values(registers))
	fmt.Println(a)
	*p1 = strconv.Itoa(a.Max)
	*p2 = strconv.Itoa(a2.Max)
}
