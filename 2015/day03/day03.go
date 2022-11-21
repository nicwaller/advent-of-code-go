package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
)

func main() {
	aoc.Select(2015, 3)
	aoc.Test(run, "sample.txt", "2", "11")
	aoc.Test(run, "sample2.txt", "", "3")
	aoc.Test(run, "input.txt", "2572", "2631")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	x := 0
	y := 0
	counts := make(map[string]int)
	for _, c := range aoc.InputString() {
		switch c {
		case '^':
			y += 1
		case 'v':
			y -= 1
		case '>':
			x += 1
		case '<':
			x -= 1
		}
		coord := fmt.Sprintf("%v-%v", x, y)
		if _, ok := counts[coord]; !ok {
			counts[coord] = 0
		}
		counts[coord]++
	}
	*p1 = strconv.Itoa(util.KeyCount(counts))

	X := make([]int, 2)
	Y := make([]int, 2)
	Counts := make(map[string]int)
	turn := 0
	homes := set.New[string]("0-0")
	for _, c := range aoc.InputString() {
		dx := 0
		dy := 0
		switch c {
		case '^':
			dy = 1
		case 'v':
			dy = -1
		case '>':
			dx = 1
		case '<':
			dx = -1
		}
		X[turn] += dx
		Y[turn] += dy
		coord := fmt.Sprintf("%v-%v", X[turn], Y[turn])
		if _, ok := Counts[coord]; !ok {
			Counts[coord] = 0
		}
		//Counts[coord]++
		homes.Insert(coord)
		turn++
		turn %= 2
	}
	*p2 = strconv.Itoa(homes.Size())
}
