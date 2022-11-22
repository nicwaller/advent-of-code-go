package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grapheasy"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2015, 13)
	aoc.Test(run, "sample.txt", "330", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLines()
	g := grapheasy.New[bool](len(lines))
	persons := set.New[string]()
	for _, line := range lines {
		//	Alice would gain 2 happiness units by sitting next to Bob.
		f := strings.Fields(line)
		name2 := strings.Trim(f[10], ".")
		cost := util.UnsafeAtoi(f[3])
		if f[2] == "lose" {
			cost = -cost
		}
		g.AddCost(f[0], name2, cost)
		persons.Extend(f[0], name2)
	}

	a1 := analyze.Box[int]{}
	for p := iter.Permute(persons.Items()); p.Next(); {
		arrangement := p.Value()
		aCost := 0
		for i, _ := range arrangement {
			aCost += g.Cost(arrangement[i], arrangement[(i+1)%len(arrangement)])
			aCost += g.Cost(arrangement[(i+1)%len(arrangement)], arrangement[i])
		}
		a1.Put(aCost)
		//fmt.Printf("%v : %6d\n", arrangement, aCost)
	}
	*p1 = strconv.Itoa(a1.Max)

	for _, pers := range persons.Items() {
		g.AddBothCost("Me", pers, 0)
	}
	persons.Add("Me")

	a2 := analyze.Box[int]{}
	for p := iter.Permute(persons.Items()); p.Next(); {
		arrangement := p.Value()
		aCost := 0
		for i, _ := range arrangement {
			aCost += g.Cost(arrangement[i], arrangement[(i+1)%len(arrangement)])
			aCost += g.Cost(arrangement[(i+1)%len(arrangement)], arrangement[i])
		}
		a2.Put(aCost)
	}
	*p2 = strconv.Itoa(a2.Max)
}
