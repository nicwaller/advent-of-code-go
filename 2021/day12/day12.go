package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grapheasy"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2021, 12)
	aoc.Test(run, "sample1.txt", "10", "36")
	aoc.Test(run, "sample2.txt", "19", "103")
	aoc.Test(run, "sample3.txt", "226", "3509")
	aoc.Test(run, "input.txt", "5333", "146553")
	aoc.Run(run)
}

type MaxNode struct {
	small bool
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLines()
	gg := grapheasy.New[MaxNode](len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "-")
		gg.AddBoth(parts[0], parts[1])
	}

	gg.DefineNodeContext(func(idx int, label string) MaxNode {
		return MaxNode{small: isSmallCave(label)}
	})

	start, _, _ := gg.NodeByName("start")

	count := 0
	getPathsFrom(start, gg, []int{start}, func(path []int) {
		count++
	}, &count)
	*p1 = strconv.Itoa(count)

	count = 0
	getPathsFrom(start, gg, []int{start}, func(path []int) {
		count++
	}, nil)
	*p2 = strconv.Itoa(count)
}

func isSmallCave(name string) bool {
	if name == "start" || name == "end" {
		return false
	}
	return strings.ToLower(name) == name
}

func getPathsFrom(v int, f grapheasy.Graph[MaxNode], pathSoFar []int, fn func(path []int), doubleCave *int) {
	isSeen := func(nodeId int) bool {
		for _, v := range pathSoFar {
			if v == nodeId {
				return true
			}
		}
		return false
	}
	f.Underlying().Visit(v, func(w int, c int64) bool {
		newPath := append(pathSoFar, w)
		_, label, ctx := f.NodeById(w)
		nodeName := *label
		if nodeName == "start" {
			return false
		}
		if nodeName == "end" {
			fn(newPath)
			return false
		}
		if isSeen(w) && ctx.small {
			if doubleCave == nil {
				getPathsFrom(w, f, newPath, fn, &w)
			}
			return false
		}
		getPathsFrom(w, f, newPath, fn, doubleCave)
		return false
	})
}
