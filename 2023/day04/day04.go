package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iterc"
	"advent-of-code/lib/set"
	"advent-of-code/lib/stack"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 4)
	aoc.Test(run, "sample.txt", "13", "30")
	aoc.Test(run, "input.txt", "25231", "9721255")
	//aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	p1sum := 0
	scoreTable := []int{0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	cardStack := stack.NewStack[int]()
	winCountList := make([]int, len(aoc.InputLines())+1)

	iterc.
		EnumerateFrom(aoc.InputLinesIterc(), 1).
		ForEach(func(card iterc.IndexedValue[string]) {
			_, numbers := util.Pair(strings.Split(card.Value, ":"))
			winPart, myPart := util.Pair(strings.Split(numbers, "|"))
			matches := set.Intersection(
				set.New(util.NumberFields(winPart)...),
				set.New(util.NumberFields(myPart)...),
			)
			winCountList[card.Index] = matches.Size()
			p1sum += scoreTable[matches.Size()]
			cardStack.Push(card.Index)
		})

	p2sum := 0
	for !cardStack.Empty() {
		card := cardStack.MustPop()
		p2sum++
		for i := 1; i <= winCountList[card]; i++ {
			cardStack.Push(card + i)
		}
	}

	*p1 = strconv.Itoa(p1sum)
	*p2 = strconv.Itoa(p2sum)
}
