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

//var currentRunPhaseConfig []int

func main() {
	aoc.Select(2019, 7)
	//currentRunPhaseConfig = []int{4, 3, 2, 1, 0}
	//aoc.Test(run, "sample1.txt", "43210", "")
	//currentRunPhaseConfig = []int{0, 1, 2, 3, 4}
	//aoc.Test(run, "sample2.txt", "54321", "")
	//currentRunPhaseConfig = []int{1, 0, 4, 3, 2}
	//aoc.Test(run, "sample3.txt", "65210", "")
	//currentRunPhaseConfig = []int{}
	aoc.Test(run, "input.txt", "67023", "7818398")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	program := f8l.Map(strings.Split(aoc.InputString(), ","), util.UnsafeAtoi)

	runPhase := func(phaseConfig []int) int {
		return ampStackRepeating(program, phaseConfig)
	}

	// this is just for testing with the samples
	//if len(currentRunPhaseConfig) > 0 {
	//	*p1 = strconv.Itoa(runPhase(currentRunPhaseConfig))
	//	return
	//}

	permutations := iter.Permute(iter.Range(0, 5).List())
	*p1 = strconv.Itoa(iter.Map(permutations, runPhase).Reduce(util.IntMax, 0))

	permutations = iter.Permute(iter.Range(5, 10).List())
	*p2 = strconv.Itoa(iter.Map(permutations, runPhase).Reduce(util.IntMax, 0))
}

func ampStackRepeating(program []int, phaseList []int) int {
	amps := make([]*intcode.IntcodeVM, 5)
	inputs := make([]chan int, 5)
	outputs := make([]chan int, 5)
	doneChan := make([]chan bool, 5)
	for i := range amps {
		inputs[i] = make(chan int, 1)
		inputs[i] <- phaseList[i]
		outputs[i] = make(chan int)
		doneChan[i] = make(chan bool)
		amps[i] = intcode.NewVM(inputs[i], outputs[i])
		amps[i].ExecAsync(program, doneChan[i])
	}
	inputs[0] <- 0
	remaining := 5
	last := make([]int, 5)
	for remaining > 0 {
		for i := range amps {
			select {
			case x := <-outputs[i]:
				last[i] = x
				inputs[(i+1)%5] <- x
			case <-doneChan[i]:
				remaining--
			}
		}
	}
	return util.Last(last)
}
