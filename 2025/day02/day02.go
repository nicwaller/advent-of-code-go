package main

import (
	"strconv"
	"strings"

	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2025, 2)
	aoc.Test(run, "sample.txt", "1227775554", "4174379265")
	aoc.Test(run, "input.txt", "23701357374", "34284458938")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	checksum1 := 0
	checksum2 := 0

	values := util.UIntFields(aoc.InputString())
	for fi := 0; fi < len(values); fi += 2 {
		// PERF: brute force: slow but simple!
		for i := values[fi]; i <= values[fi+1]; i++ {
			s := strconv.Itoa(i)
			if isRepeatedOnce(s) {
				checksum1 += i
			}
			if isRepeatingSequence(s) {
				checksum2 += i
			}
		}
	}
	*p1 = strconv.Itoa(checksum1)
	*p2 = strconv.Itoa(checksum2)
}

func isRepeatedOnce(s string) bool {
	if len(s)%2 == 1 {
		// can never be true for odd length strings
		return false
	}
	b := len(s) / 2
	h1 := s[:b]
	h2 := s[b:]
	return h1 == h2
}

func isRepeatingSequence(s string) bool {
	maxLen := len(s) / 2
	for patternLength := 1; patternLength <= maxLen; patternLength++ {
		if isRepeatingSequenceN(s, patternLength) {
			return true
		}
	}
	return false
}

func isRepeatingSequenceN(haystack string, patternLength int) bool {
	if len(haystack)%patternLength != 0 {
		return false
	}
	needle := haystack[0:patternLength]
	patternReps := len(haystack) / patternLength
	return haystack == strings.Repeat(needle, patternReps)
}
