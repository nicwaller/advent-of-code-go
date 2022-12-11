package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/queue"
	"advent-of-code/lib/util"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 11)
	aoc.Test(run, "sample.txt", "10605", "2713310158")
	//aoc.Test(run, "input.txt", "55458", "14508081294")
	aoc.Run(run)
}

type Monkey struct {
	inventory queue.Queue[int]
	operation func(v int) int
	testMod   int
	targetSel map[bool]int
}

func parse() []*Monkey {
	monkeys := make([]*Monkey, 0)
	for _, lines := range aoc.ParagraphsIterator().List() {
		m := Monkey{
			inventory: queue.FromSlice[int](util.NumberFields(lines[1])),
			testMod:   util.NumberFields(lines[3])[0],
			targetSel: map[bool]int{
				true:  util.NumberFields(lines[4])[0],
				false: util.NumberFields(lines[5])[0],
			},
		}
		operator := strings.Fields(lines[2])[4]
		operand := strings.Fields(lines[2])[5]
		switch {
		case operator == "*" && operand == "old":
			m.operation = func(v int) int {
				return v * v
			}
		case operator == "+":
			operandN := util.UnsafeAtoi(operand)
			m.operation = func(v int) int {
				return v + operandN
			}
		case operator == "*":
			operandN := util.UnsafeAtoi(operand)
			m.operation = func(v int) int {
				return v * operandN
			}
		default:
			panic(operator)
		}
		monkeys = append(monkeys, &m)
	}
	return monkeys
}

func simSimian(rounds int, op func(int) int) int {
	monkeys := parse()
	monkeyInspections := make([]int, len(monkeys))
	for round := 1; round <= rounds; round++ {
		for monkeyN, m := range monkeys {
			for m.inventory.Length() > 0 {
				monkeyInspections[monkeyN]++
				itemScore := op(m.operation(m.inventory.MustPop()))
				monkeys[m.targetSel[itemScore%m.testMod == 0]].inventory.Push(itemScore)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyInspections)))
	return util.IntProductV(monkeyInspections[0:2]...)
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(simSimian(20, func(v int) int { return v / 3 }))
	magicMod := f8l.Reduce(f8l.Map(parse(), func(m *Monkey) int { return m.testMod }), 1, util.IntProduct)
	*p2 = strconv.Itoa(simSimian(10000, func(v int) int { return v % magicMod }))
}
