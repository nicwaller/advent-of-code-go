package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Day(2)
	aoc.Test(run, "sample.txt", "150", "900")
	aoc.Test(run, "input.txt", "1580000", "1251263225")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	var y, z = 0, 0
	var z2, aim = 0, 0
	for lines := aoc.InputLinesIterator(); lines.Next(); {
		fields := strings.Split(lines.Value(), " ")
		direction := fields[0]
		quantity := util.UnsafeAtoi(fields[1])
		switch direction {
		case "up":
			z -= quantity
			aim -= quantity
		case "down":
			z += quantity
			aim += quantity
		case "forward":
			y += quantity
			z2 += quantity * aim
		}
	}
	*p1 = strconv.Itoa(y * z)
	*p2 = strconv.Itoa(y * z2)
}
