package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/f8l"
	"fmt"
	"sort"
	"strings"
)

func main() {
	aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

func parseFile() []string {
	return aoc.InputLines()
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

func score(badChar rune) int {
	switch badChar {
	case ' ':
		return 0
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		panic(badChar)
	}
}

func autoScore(chars []rune) int {
	iscore := 0
	for _, c := range chars {
		iscore *= 5
		switch c {
		case ')':
			iscore += 1
		case ']':
			iscore += 2
		case '}':
			iscore += 3
		case '>':
			iscore += 4
		default:
			panic(c)
		}
	}
	return iscore
}

func part1(input []string) int {
	errs := f8l.Map[string, rune](&input, firstSyntaxError)
	//fmt.Println(errs)
	scores := f8l.Map[rune, int](&errs, score)
	//fmt.Println(scores)
	totalErrorScore := f8l.Reduce[int](&scores, 0, func(a int, b int) int { return a + b })
	assert.EqualAny(totalErrorScore, []int{26397, 265527}, "total error score")
	return totalErrorScore
}

func part2(input []string) int {
	// Now, discard the corrupted lines. The remaining lines are incomplete.
	lines := f8l.Filter(&input, func(line string) bool { return ' ' == firstSyntaxError(line) })
	scores := f8l.Map[string, int](&lines, func(line string) int {
		return autoScore(autocomplete(line))
	})
	sort.Ints(scores)
	middleScore := scores[len(scores)/2]
	assert.EqualAny(middleScore, []int{288957, 3969823589}, "middle score")
	return middleScore
}
