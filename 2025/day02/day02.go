package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"

	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2025, 2)
	aoc.Test(run, "sample.txt", "1227775554", "")
	aoc.Test(run, "input.txt", "23701357374", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	var checksum1 uint64 = 0
	for _, line := range aoc.InputLines() {
		ranges := strings.Split(line, ",")
		for _, strRange := range ranges {
			if strRange == "" {
				continue
			}
			elements := strings.Split(strRange, "-")
			first := util.UnsafeAtoi(elements[0])
			last := util.UnsafeAtoi(elements[1])
			for iid := range invalidIDsInRange(uint64(first), uint64(last)) {
				checksum1 += iid
			}
		}
	}
	*p1 = strconv.Itoa(int(checksum1))
}

func invalidIDsInRange(first, last uint64) iter.Seq[uint64] {
	if last < first {
		panic("invalid range")
	}
	return func(yield func(uint64) bool) {
		strFirst := fmt.Sprintf("%d", first)
		strFirstHalf := strFirst[:(len(strFirst) / 2)]
		v, _ := strconv.Atoi(strFirstHalf)
		var ticker = uint64(v)
		var tickerConcat uint64
		update := func(increment bool) {
			if increment {
				ticker++
			}
			tickerConcat = uint64(util.UnsafeAtoi(fmt.Sprintf("%d%d", ticker, ticker)))
		}

		update(false)
		for ; tickerConcat <= last; update(true) {
			if tickerConcat >= first && tickerConcat <= last {
				yield(tickerConcat)
			} else {
				// wasted effort
			}
		}
	}
}
