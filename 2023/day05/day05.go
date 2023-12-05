package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iterc"
	"advent-of-code/lib/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 5)
	aoc.Test(run, "sample.txt", "35", "46")
	aoc.Test(run, "input.txt", "825516882", "136096660")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {

	pgs := aoc.InputParagraphs()
	//seeds := util.NumberFields(pgs[0][0])
	seeds := strings.Fields(pgs[0][0])[1:]
	pgs = pgs[1:]

	fmt.Println("seeds:", seeds)

	maps := make([][]func(uint64) (uint64, bool), 0)
	for _, pg := range pgs {
		//ff := strings.Fields(pg[0])
		//mapName := ff[0]
		mapsInner := make([]func(uint64) (uint64, bool), 0)
		for i := 1; i < len(pg); i++ {
			fff := util.NumberFields(pg[i])
			srcRangeStart := uint64(fff[1])
			destRangeStart := uint64(fff[0])
			rangeLength := uint64(fff[2])
			mapsInner = append(mapsInner, func(v uint64) (uint64, bool) {
				if v >= srcRangeStart && v <= srcRangeStart+rangeLength {
					return destRangeStart + (v - srcRangeStart), true
				} else {
					return v, false
				}
			})
		}
		maps = append(maps, mapsInner)
	}

	var found bool
	var lowest1 uint64 = math.MaxInt64
	for _, seed := range seeds {
		v := uint64(util.UnsafeAtoi(seed))
		for _, m := range maps {
		transforms:
			for _, transform := range m {
				if v, found = transform(v); found {
					break transforms
				}
			}
			//if !found {
			//	v =v
			//}
		}
		lowest1 = min(lowest1, v)
		fmt.Println("seed", seed, v)
	}

	var lowest2 uint64 = math.MaxInt64
	iterc.Chunk(iterc.ListIterator(seeds), 2).ForEach(func(w []string) {
		if len(w) == 0 {
			return
		}
		offset := uint64(util.UnsafeAtoi(w[0]))
		rangelen := uint64(util.UnsafeAtoi(w[1]))
		terminus := offset + rangelen
		var i uint64
		for i = offset; i < terminus; i++ {
			if i%10000000 == 0 {
				fmt.Println("10million") // 20,000 of these
			}
			v := i
			for _, m := range maps {
			transforms:
				for _, transform := range m {
					if v, found = transform(v); found {
						break transforms
					}
				}
			}
			lowest2 = min(lowest2, v)
		}
	})

	*p1 = strconv.Itoa(int(lowest1))
	*p2 = strconv.Itoa(int(lowest2))
}

//func (d *DefaultMap) Get(k string) string {
//	for _, f := range *d {
//		if v, found := f(k); found {
//			return v
//		}
//	}
//	return k
//}

func resolve(seed string, mm []func(string) (string, bool)) {

}
