package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/aoc/intcode"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

var currentRunPhaseConfig []int

func main() {
	aoc.Select(2019, 7)
	currentRunPhaseConfig = []int{4, 3, 2, 1, 0}
	aoc.Test(run, "sample1.txt", "43210", "")
	currentRunPhaseConfig = []int{0, 1, 2, 3, 4}
	aoc.Test(run, "sample2.txt", "54321", "")
	currentRunPhaseConfig = []int{1, 0, 4, 3, 2}
	aoc.Test(run, "sample3.txt", "65210", "")
	currentRunPhaseConfig = []int{}
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	program := f8l.Map(strings.Split(aoc.InputString(), ","), util.UnsafeAtoi)

	ampStack := func(phaseList []int) int {
		var last int
		for _, phase := range phaseList {
			result := intcode.ExecIO(program, []int{phase, last})
			last = util.Last(result)
		}
		return last
	}

	// this is just for testing with the samples
	if len(currentRunPhaseConfig) > 0 {
		*p1 = strconv.Itoa(ampStack(currentRunPhaseConfig))
		return
	}

	permutations := iter.Permute(iter.Range(0, 5).List())
	thrusts := iter.Map[[]int, int](permutations, func(perm []int) int {
		return ampStack(perm)
	})
	maxThrust := thrusts.Reduce(util.IntMax, 0)
	*p1 = strconv.Itoa(maxThrust)
	*p2 = strconv.Itoa(0)
}
