package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grapheasy"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2019, 6)
	aoc.Test(run, "sample.txt", "42", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLines()
	g := grapheasy.New[bool](len(lines) * 2)
	for _, line := range lines {
		c, o := util.Pair(strings.Split(line, ")"))
		g.Add(c, o)
	}

	start, _, _ := g.NodeByName("COM")
	count := 0
	g.DFS(start, func(v int, label string, _ *bool, path []int) {
		count += len(path)
	})
	*p1 = strconv.Itoa(count)

	*p2 = strconv.Itoa(0)
}
