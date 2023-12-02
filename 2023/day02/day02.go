package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 2)
	aoc.Test(run, "sample.txt", "8", "2286")
	aoc.Test(run, "input.txt", "2593", "54699")
	//aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	const (
		r = "red"
		g = "green"
		b = "blue"
	)

	const rMax = 12
	const gMax = 13
	const bMax = 14

	sum := 0
	powerSum := 0
	for i, line := range aoc.InputLines() {
		include := true
		gameId := i + 1
		_, gameData := util.Pair(strings.Split(line, ": "))
		rounds := strings.Split(gameData, "; ")
		cubeMax := map[string]int{r: 0, g: 0, b: 0}
		for _, round := range rounds {
			cubeCount := map[string]int{r: 0, g: 0, b: 0}
			colours := strings.Split(round, ", ")
			for _, c := range colours {
				countS, colour := util.Pair(strings.Fields(c))
				count := util.UnsafeAtoi(countS)
				cubeCount[colour] += count
				// TODO: use a newer version of Go for built-in max() function
				cubeMax[colour] = util.IntMax(count, cubeMax[colour])
			}
			if cubeCount[r] > rMax || cubeCount[g] > gMax || cubeCount[b] > bMax {
				include = false
			}
		}
		// part 1
		if include {
			sum += gameId
		}
		// part 2
		power := cubeMax[r] * cubeMax[g] * cubeMax[b]
		powerSum += power
	}

	*p1 = strconv.Itoa(sum)
	*p2 = strconv.Itoa(powerSum)
}
