package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 6)
	aoc.Test(run, "sample.txt", "288", "71503")
	aoc.Test(run, "input.txt", "2756160", "34788142")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLines()
	raceDurations := util.NumberFields(lines[0])
	distanceRecords := util.NumberFields(lines[1])
	product := 1
	for raceId := 0; raceId < len(raceDurations); raceId++ {
		raceDuration := raceDurations[raceId]
		raceRecord := distanceRecords[raceId]
		product *= winWays(raceDuration, raceRecord)
	}

	bigRaceDuration := fuzzyNumber(lines[0])
	bigRaceRecord := fuzzyNumber(lines[1])
	bigWinWays := winWays(bigRaceDuration, bigRaceRecord)

	*p1 = strconv.Itoa(product)
	*p2 = strconv.Itoa(bigWinWays)
}

func fuzzyNumber(s string) int {
	var intMatcher = regexp.MustCompile("[0-9]")
	matches := intMatcher.FindAllString(s, math.MaxInt32)
	return util.UnsafeAtoi(strings.Join(matches, ""))
}

func winWays(raceDuration int, raceRecord int) int {
	return iter.Range(0, raceDuration).
		Map(distanceCalcer(raceDuration)).
		Filter(func(i int) bool { return i > raceRecord }).
		Count()
	// PERF: the iterc package is 60x slower than iter. fuck.

	//possibleAccelTimes := iterc.Range(0, raceDuration)
	//return iterc.
	//	Map[int, int](possibleAccelTimes, distanceCalcer(raceDuration)).
	//	Filter(func(i int) bool { return i > raceRecord }).
	//	Count()
}

func distanceCalcer(raceTime int) func(int) int {
	return func(accelTime int) int {
		if accelTime < 0 {
			panic(accelTime)
		}
		if accelTime > raceTime {
			panic(accelTime)
		}
		goTime := raceTime - accelTime
		return accelTime * goTime
	}
}
