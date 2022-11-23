package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2015, 5)
	aoc.TestLiteral(run, "ugknbfddgicrmopn", "1", "")
	aoc.TestLiteral(run, "aaa", "1", "")
	aoc.TestLiteral(run, "jchzalrnumimnmhp", "0", "")
	aoc.TestLiteral(run, "haegwjzuvuyypxyu", "0", "")
	aoc.TestLiteral(run, "dvszwmarrgswjxmb", "0", "")
	aoc.TestLiteral(run, "qjhvhtzxzqqjkmpb", "", "1")
	aoc.TestLiteral(run, "xxyxx", "", "1")
	aoc.TestLiteral(run, "uurcxstgmygtbstg", "", "0")
	aoc.TestLiteral(run, "ieodomkazucvgmuy", "", "0")
	aoc.TestLiteral(run, "joakcwpfggtitizs", "", "1")
	aoc.Test(run, "input.txt", "238", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(aoc.InputLinesIterator().Filter(isNice1).Count())
	*p2 = strconv.Itoa(aoc.InputLinesIterator().Filter(isNice2).Echo().Count())
}

func isNice1(s string) bool {
	vowels := iter.StringIterator(s, 1).Filter(func(s string) bool {
		return strings.Contains("aeiou", s)
	}).Count()
	//distinctVowels := analyze.Analyze(vowels).Distinct
	doubles := iter.SlidingWindow(2, iter.StringIterator(s, 1)).Filter(func(s []string) bool {
		return s[0] == s[1]
	}).Count()
	rejection := strings.Contains(s, "ab") ||
		strings.Contains(s, "cd") ||
		strings.Contains(s, "pq") ||
		strings.Contains(s, "xy")
	return vowels >= 3 && doubles >= 1 && !rejection
}

func isNice2(s string) bool {
	hasGapLetter := iter.SlidingWindow(3, iter.StringIterator(s, 1)).
		Filter(func(trip []string) bool {
			return trip[0] == trip[2]
		}).Count() > 0
	if !hasGapLetter {
		return false
	}

	gotOne := false
	for pairs := iter.SlidingWindow[string](2, iter.StringIterator(s, 1)); pairs.Next(); {
		pair := strings.Join(pairs.Value(), "")
		if strings.Count(s, pair) < 2 {
			continue
		}
		gotOne = true
	}

	return gotOne
}
