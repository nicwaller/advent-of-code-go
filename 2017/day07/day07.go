package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grapheasy"
	"advent-of-code/lib/util"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2017, 7)
	aoc.Test(run, "sample.txt", "tknk", "60")
	//aoc.Test(run, "input.txt", "vvsvez", "362")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLines()
	g := grapheasy.New[int](len(lines) * 2)
	for _, line := range lines {
		ff := strings.Fields(line)
		name := ff[0]
		weight := util.NumberFields(ff[1])[0]
		g.AddContext(name, weight)
		if len(ff) > 2 {
			children := ff[3:]
			for _, child := range children {
				childName := strings.TrimSuffix(child, ",")
				g.Add(name, childName)
			}
		}
	}
	headId, headName, _ := g.Head()
	*p1 = *headName

	goodWeight := 0
	findError(headId, g, false, &goodWeight)
	*p2 = strconv.Itoa(goodWeight)
}

func findError(v int, g grapheasy.Graph[int], engage bool, goodWeight *int) bool {
	childNodes := make([]int, 0)
	childWeights := make([]int, 0)
	totalWeights := make([]int, 0)
	g.Underlying().Visit(v, func(w int, c int64) bool {
		childNodes = append(childNodes, w)
		_, _, childWeight := g.NodeById(w)
		childWeights = append(childWeights, *childWeight)
		totalWeights = append(totalWeights, totalWeight(w, g))
		return false
	})
	a := analyze.Analyze(totalWeights)
	diff := a.Max - a.Min
	if diff == 0 && !engage {
		// found a normal healthy balance
		return false
	} else if diff == 0 && engage {
		// found a healthy balance inside the lopsided branch
		return false
	} else if diff > 0 {
		// found a sign of trouble!
		// identify the troublesome sub-branch:
		idx := slices.Index(totalWeights, a.LeastCommon)
		// and keep delving into it
		b := findError(childNodes[idx], g, true, goodWeight)
		isProblemChild := !b
		if isProblemChild {
			badWeight := childWeights[idx]
			*goodWeight = badWeight - diff
		}
		return true
	}
	return false
}

func totalWeight(fromNode int, g grapheasy.Graph[int]) int {
	sum := 0
	g.DFS(fromNode, func(v int, label string, ctx *int, path []int) {
		sum += *ctx
	})
	return sum
}
