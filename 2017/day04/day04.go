package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2017, 4)
	aoc.Test(run, "sample.txt", "", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aa := iter.StringIterator(a, 1).List()
	bb := iter.StringIterator(b, 1).List()
	sort.Strings(aa)
	sort.Strings(bb)
	for i := 0; i < len(aa); i++ {
		if aa[i] != bb[i] {
			return false
		}
	}
	return true
}

func run(p1 *string, p2 *string) {
	valid := 0
	for _, passphrase := range aoc.InputLines() {
		words := strings.Fields(passphrase)
		uniqueWords := set.FromSlice(words)
		if len(words) == uniqueWords.Size() {
			valid++
		}
	}
	*p1 = strconv.Itoa(valid)

	valid2 := 0
checkLoop:
	for _, passphrase := range aoc.InputLines() {
		words := strings.Fields(passphrase)
		pairs := iter.CombinationsN(words, 2)
		for pairs.Next() {
			p := pairs.Value()
			if isAnagram(p[0], p[1]) {
				continue checkLoop
			}
		}
		valid2++
	}

	*p2 = strconv.Itoa(valid2)
}
