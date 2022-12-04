package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"strconv"
)

func main() {
	aoc.Select(2022, 3)
	aoc.Test(run, "sample.txt", "157", "70")
	aoc.Test(run, "input.txt", "8088", "")
	aoc.Run(run)
}

func itemPriority(s string) uint8 {
	isLowercase := (s[0] & 0b00100000) != 0
	inPlaneVal := s[0] & 0b00011111
	if isLowercase {
		return inPlaneVal
	} else {
		return inPlaneVal + 26
	}
}

func run(p1 *string, p2 *string) {
	pSum := 0
	for _, packingList := range aoc.InputLines() {
		rucksack1Str := packingList[0 : len(packingList)/2]
		rucksack2Str := packingList[len(packingList)/2:]
		rucksack1 := set.FromIterable(iter.StringIterator(rucksack1Str, 1))
		rucksack2 := set.FromIterable(iter.StringIterator(rucksack2Str, 1))
		ix := set.Intersection(rucksack1, rucksack2)
		pSum += int(itemPriority(ix.Items()[0]))
	}
	*p1 = strconv.Itoa(pSum)

	groupIter := iter.Chunk[string](3, aoc.InputLinesIterator())
	pSum2 := 0
	for groupIter.Next() {
		group := groupIter.Value()
		ix2 := set.Intersection(
			set.FromIterable(iter.StringIterator(group[0], 1)),
			set.FromIterable(iter.StringIterator(group[1], 1)),
			set.FromIterable(iter.StringIterator(group[2], 1)),
		)
		pSum2 += int(itemPriority(ix2.Items()[0]))
	}
	*p2 = strconv.Itoa(pSum2)
}
