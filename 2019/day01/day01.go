package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
)

func main() {
	aoc.Select(2019, 1)
	aoc.Test(run, "sample.txt", "33583", "50346")
	aoc.Test(run, "input.txt", "3216744", "4822249")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	totalFuel := f8l.Reduce(f8l.Map(aoc.InputLinesInt(), fuelReq),
		0,
		util.IntSum)
	*p1 = strconv.Itoa(totalFuel)

	totalFuel2 := f8l.Reduce(f8l.Map(aoc.InputLinesInt(), fuelRecursive),
		0,
		util.IntSum)
	*p2 = strconv.Itoa(totalFuel2)
}

func fuelReq(mass int) int {
	return mass/3 - 2
}

func fuelRecursive(mass int) int {
	moreFuel := util.IntMax(0, fuelReq(mass))
	fmt.Printf("%d -> %d\n", mass, moreFuel)
	if moreFuel == 0 {
		return 0
	}
	return moreFuel + fuelRecursive(moreFuel)
}
