package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/f8l"
	"fmt"
	"strings"
)

func main() {
	//aoc.UseSampleFile()
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
	for i, c := range line {
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
				fmt.Printf("error in col %d (%v != %v)\n", i, string(actual), string(expected))
				return c
			}
		}
	}
	return ' '
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

func part1(input []string) int {
	errs := f8l.Map[string, rune](&input, firstSyntaxError)
	fmt.Println(errs)
	scores := f8l.Map[rune, int](&errs, score)
	fmt.Println(scores)
	totalErrorScore := f8l.Reduce[int](&scores, 0, func(a int, b int) int { return a + b })
	assert.Equal(totalErrorScore, 265527)
	return totalErrorScore
}

func part2(input []string) int {
	//assert.NotEqual(totalFish, 5934)
	//assert.EqualAny(best, []int{168, 97038163}, "best")
	return -1
}
