package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grapheasy"
	"fmt"
	"github.com/yourbasic/graph"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2017, 12)
	aoc.Test(run, "sample.txt", "6", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLines()
	g := grapheasy.New[string](len(lines))
	for _, line := range lines {
		ff := strings.Fields(line)
		a := ff[0]
		for _, link := range ff[2:] {
			b := strings.TrimSuffix(link, ",")
			g.Add(a, b)
		}
	}
	connected := 0
	zeroNode, _, _ := g.NodeByName("0")

	components := graph.Components(g.Underlying())
	for _, ss := range components {
		if slices.Index(ss, zeroNode) != -1 {
			fmt.Println(ss)
			connected = len(ss)
		}
	}

	*p1 = strconv.Itoa(connected)
	*p2 = strconv.Itoa(len(components))
}
