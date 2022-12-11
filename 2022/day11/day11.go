package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/queue"
	"advent-of-code/lib/util"
	"fmt"
	"math/big"
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
	inventory queue.Queue[*big.Int]
	operation func(v *big.Int)
	testMod   int
	targetSel map[bool]int
}

func parse() []*Monkey {
	monkeys := make([]*Monkey, 0)
	for _, lines := range aoc.ParagraphsIterator().List() {
		bigNf := f8l.Map[int, *big.Int](util.NumberFields(lines[1]), func(v int) *big.Int {
			return big.NewInt(int64(v))
		})
		m := Monkey{
			inventory: queue.FromSlice[*big.Int](bigNf),
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
			m.operation = func(v *big.Int) {
				if operand == -1 {
					v.Add(v, v)
				} else {
					v.Add(v, big.NewInt(int64(operand)))
				}
			}
		case "*":
			m.operation = func(v *big.Int) {
				if operand == -1 {
					v.Mul(v, v)
				} else {
					v.Mul(v, big.NewInt(int64(operand)))
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
			var itemScore *big.Int
			for m.inventory.Length() > 0 {
				monkeyInspections[monkeyN]++
				itemScore, _ = m.inventory.Pop()
				//fmt.Printf(" Monkey inspects an item with a worry level of %d:\n", itemScore)
				m.operation(itemScore)
				//fmt.Printf("  Worry level becomes %d:\n", itemScore)
				itemScore.Div(itemScore, big.NewInt(3))
				//fmt.Printf("  Monkey gets bored with item. Worry level is divided by 3 to %d\n", itemScore)
				modResult := big.NewInt(0)
				modResult.Set(itemScore)
				modResult.Mod(modResult, big.NewInt(int64(m.testMod)))
				sel := modResult.Int64() == 0
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
			var itemScore *big.Int
			for m.inventory.Length() > 0 {
				monkeyInspections[monkeyN]++
				itemScore, _ = m.inventory.Pop()
				if debug {
					fmt.Printf(" Monkey inspects an item with a worry level of %d:\n", itemScore)
				}

				m.operation(itemScore)
				if debug {
					fmt.Printf("  Worry level becomes %d:\n", itemScore)
				}

				itemScore.Mod(itemScore, big.NewInt(int64(magicMod)))
				if debug {
					fmt.Printf("  Monkey gets bored with item. Worry level is reduced to %d\n", itemScore)
				}

				sel := itemScore.Int64()%int64(m.testMod) == 0
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
