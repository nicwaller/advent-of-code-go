package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc.Day(10)
	aoc.Test(run, "sample.txt", "26397", "288957")
	aoc.Test(run, "input.txt", "265527", "3969823589")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(part1())
	*p2 = strconv.Itoa(part2())
}

func firstSyntaxError(line string) rune {
	match := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	deque := make([]rune, 100)
	dPtr := -1
	for _, c := range line {
		if strings.ContainsAny(string(c), "([{<") {
			// opening match
			dPtr++
			deque[dPtr] = c
		} else {
			// should be closing match
			actual := c
			expected := match[deque[dPtr]]
			if actual == expected {
				dPtr--
			} else {
				//fmt.Printf("error in col %d (%v != %v)\n", i, string(actual), string(expected))
				return c
			}
		}
	}
	return ' '
}

func autocomplete(line string) []rune {
	match := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	deque := make([]rune, 100)
	dPtr := -1
	for _, c := range line {
		if strings.ContainsAny(string(c), "([{<") {
			// opening match
			dPtr++
			deque[dPtr] = match[c]
		} else {
			// should be closing match
			actual := c
			expected := deque[dPtr]
			if actual == expected {
				dPtr--
			} else {
				panic("this should not be processing corrupt lines")
			}
		}
	}
	reverse(deque[:dPtr+1])
	return deque[:dPtr+1]
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

var errScoreTable = map[rune]int{
	' ': 0,
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func score(badChar rune) int {
	return errScoreTable[badChar]
}

var acScoreTable = map[rune]int{
	' ': 0,
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func autoScore(chars []rune) int {
	iscore := 0
	for _, c := range chars {
		iscore = (iscore * 5) + acScoreTable[c]
	}
	return iscore
}

func part1() int {
	input := aoc.InputLines()
	errs := f8l.Map[string, rune](input, firstSyntaxError)
	scores := f8l.Map[rune, int](errs, score)
	totalErrorScore := f8l.Reduce[int](scores, 0, util.IntSum)
	return totalErrorScore
}

func part2() int {
	input := aoc.InputLines()
	// Now, discard the corrupted lines. The remaining lines are incomplete.
	lines := f8l.Filter(&input, func(line string) bool { return ' ' == firstSyntaxError(line) })
	scores := f8l.Map[string, int](lines, func(line string) int {
		return autoScore(autocomplete(line))
	})
	sort.Ints(scores)
	middleScore := scores[len(scores)/2]
	return middleScore
}
