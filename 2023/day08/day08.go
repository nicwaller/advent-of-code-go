package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"golang.org/x/exp/maps"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 8)
	aoc.Test(run, "sample.txt", "2", "")
	aoc.Test(run, "sample2.txt", "6", "")
	aoc.Test(run, "sample3.txt", "", "6")
	aoc.Test(run, "input.txt", "19951", "16342438708751")
	aoc.Run(run)
	aoc.Out()
}

type Place struct {
	name     string
	left     string
	right    string
	shortcut [3]string
}

func run(p1 *string, p2 *string) {
	pathPg, graphPg := util.Pair(aoc.InputParagraphs())
	places := map[string]Place{}
	for _, line := range graphPg {
		ff := util.AlphanumericFields(line)
		places[ff[0]] = Place{
			name:  ff[0],
			left:  ff[1],
			right: ff[2],
			shortcut: [3]string{
				"", ff[1], ff[2],
			},
		}
	}
	path := pathPg[0]
	cur := "AAA"
	c := 0
	if _, found := places[cur]; !found {
		// skip part 1 if the sample is not suitable
		goto part2
	}
	iter.StringIterator(path, 1).Repeat().TakeWhile(func(_ string) bool {
		return cur != "ZZZ"
	}).Counting(&c).Each(func(choice string) {
		cur = places[cur].shortcut[rune(choice[0])%5]
	})
	*p1 = strconv.Itoa(c)

part2:
	states := iter.ListIterator(maps.Keys(places)).Filter(func(s string) bool {
		return strings.HasSuffix(s, "A")
	}).List()
	if len(states) < 2 {
		// skip part 2 if the sample is not suitable
		return
	}
	cycleLen := make([]int, len(states))
	for i := range states {
		z := 0
		keepGoing := true
		for ite := iter.StringIterator(path, 1).Repeat(); ite.Next() && keepGoing; z++ {
			choice := ite.Value()
			states[i] = places[states[i]].shortcut[rune(choice[0])%5]
			keepGoing = states[i][2] != 'Z'
		}
		cycleLen[i] = z
	}
	*p2 = strconv.Itoa(util.LCM_V(cycleLen...))
}
