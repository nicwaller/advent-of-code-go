package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/util"
	"fmt"
)

func main() {
	//aoc.UseSampleFile()
	fmt.Printf("Part 1: %d\n", part1(parseFile()))
	fmt.Printf("Part 2: %d\n", part2(parseFile()))
}

type fileType [2]int

func parseFile() fileType {
	lines := aoc.InputLines()
	return [2]int{
		util.UnsafeAtoi(lines[0][len(lines[0])-1:]),
		util.UnsafeAtoi(lines[1][len(lines[1])-1:]),
	}
}

func part1(input fileType) int {
	fmt.Println(input)

	// prepare track
	//track := make([]int, 10)

	// deterministic dice
	lastRoll := -1
	totalRolls := 0
	rollDie := func() int {
		totalRolls++
		lastRoll = (lastRoll + 1) % 100
		return lastRoll + 1
	}

	p1Pos := input[0] - 1
	p2Pos := input[1] - 1
	p1Score := 0
	p2Score := 0
	for {
		rollTotal := rollDie() + rollDie() + rollDie()
		p1Pos = (p1Pos + rollTotal) % 10
		p1Score += p1Pos + 1
		fmt.Printf("Player 1 rolls ... and moves to space %d for a total score of %d.\n", p1Pos+1, p1Score)
		if p1Score >= 1000 {
			fmt.Printf("Total die rolls: %d\n", totalRolls)
			fmt.Printf("P1: %d\n", p1Score)
			fmt.Printf("P2: %d\n", p2Score)
			assert.EqualAny(p2Score*totalRolls, []int{908091}, "part 1")
			return p2Score * totalRolls
		}

		rollTotal = rollDie() + rollDie() + rollDie()
		p2Pos = (p2Pos + rollTotal) % 10
		p2Score += p2Pos + 1
		fmt.Printf("Player 2 rolls ... and moves to space %d for a total score of %d.\n", p2Pos+1, p2Score)
		if p2Score >= 1000 {
			fmt.Printf("Total die rolls: %d\n", totalRolls)
			fmt.Printf("P1: %d\n", p1Score)
			fmt.Printf("P2: %d\n", p2Score)
			return p1Score * totalRolls
		}
	}
}

func part2(input fileType) int {
	return -1
}
