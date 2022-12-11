package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/queue"
	"advent-of-code/lib/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 11)
	aoc.Test(run, "sample.txt", "10605", "2713310158")
	//aoc.Test(run, "input.txt", "55458", "14508081294")
	aoc.Run(run)
	aoc.Out()
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
		nf := util.NumberFields(lines[2])
		var operand int
		switch len(nf) {
		case 0:
			operand = -1
		case 1:
			operand = nf[0]
		default:
			panic(nf)
		}
		switch operator {
		case "+":
			m.operation = func(v int) int {
				if operand == -1 {
					return v + v
				} else {
					return v + operand
				}
			}
		case "*":
			m.operation = func(v int) int {
				if operand == -1 {
					return v * v
				} else {
					return v * operand
				}
			}
		default:
			panic(operator)
		}
		monkeys = append(monkeys, &m)
	}
	return monkeys
}

//goland:noinspection GoBoolExpressions
func run(p1 *string, p2 *string) {
	monkeyInspections := make([]int, 8)
	monkeys := parse()
	for round := 1; round <= 20; round++ {
		for monkeyN, m := range monkeys {
			//fmt.Printf("Monkey %d:\n", monkeyN)
			if m.inventory.Length() == 0 {
				continue
			}
			var itemScore int
			for m.inventory.Length() > 0 {
				monkeyInspections[monkeyN]++
				itemScore, _ = m.inventory.Pop()
				//fmt.Printf(" Monkey inspects an item with a worry level of %d:\n", itemScore)
				itemScore = m.operation(itemScore)
				//fmt.Printf("  Worry level becomes %d:\n", itemScore)
				itemScore /= 3
				//fmt.Printf("  Monkey gets bored with item. Worry level is divided by 3 to %d\n", itemScore)
				sel := itemScore%m.testMod == 0
				//fmt.Printf("  Current worry level divisible by 23? %v\n", sel)
				target := m.targetSel[sel]
				//fmt.Printf("  Item with worry level %d is thrown to monkey %d\n", itemScore, target)
				err := monkeys[target].inventory.Push(itemScore)
				if err != nil {
					panic(err)
				}
			}
		}
		//fmt.Printf("After Round %d:\n", round)
		//for monkeyNum, m := range monkeys {
		//	fmt.Printf(" Monkey %d: %v\n", monkeyNum, m.inventory.Items())
		//}
	}
	sort.Ints(monkeyInspections)
	fmt.Println(monkeyInspections)
	ll := len(monkeyInspections)
	*p1 = strconv.Itoa(monkeyInspections[ll-1] * monkeyInspections[ll-2])

	testMods := f8l.Map[*Monkey, int](monkeys, func(m *Monkey) int { return m.testMod })
	magicMod := f8l.Reduce(testMods, 1, util.IntProduct)

	debug := false

	monkeyInspections = make([]int, 8)
	monkeys = parse()
	for round := 1; round <= 10000; round++ {
		if debug {
			fmt.Printf("Round %d\n", round)
		}
		for monkeyN, m := range monkeys {
			if debug {
				fmt.Printf("Monkey %d:\n", monkeyN)
			}
			if m.inventory.Length() == 0 {
				continue
			}
			var itemScore int
			for m.inventory.Length() > 0 {
				monkeyInspections[monkeyN]++
				itemScore, _ = m.inventory.Pop()
				if debug {
					fmt.Printf(" Monkey inspects an item with a worry level of %d:\n", itemScore)
				}

				itemScore = m.operation(itemScore)
				if debug {
					fmt.Printf("  Worry level becomes %d:\n", itemScore)
				}

				itemScore %= magicMod
				if debug {
					fmt.Printf("  Monkey gets bored with item. Worry level is reduced to %d\n", itemScore)
				}

				sel := itemScore%m.testMod == 0
				//fmt.Printf("  Current worry level divisible by 23? %v\n", sel)
				target := m.targetSel[sel]
				//fmt.Printf("  Item with worry level %d is thrown to monkey %d\n", itemScore, target)
				err := monkeys[target].inventory.Push(itemScore)
				if err != nil {
					panic(err)
				}
			}
		}
		if round%1000 == 0 || round == 1 || round == 20 {
			fmt.Printf("After Round %d:\n", round)
			for id, v := range monkeyInspections[0:4] {
				fmt.Printf(" Monkey %d inspected items: %v\n", id, v)
			}
		}
	}
	sort.Ints(monkeyInspections)
	fmt.Println(monkeyInspections)
	ll = len(monkeyInspections)
	*p2 = strconv.Itoa(monkeyInspections[ll-1] * monkeyInspections[ll-2])
}
