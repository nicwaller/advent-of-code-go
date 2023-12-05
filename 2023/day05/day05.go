package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 5)
	aoc.Test(run, "sample.txt", "35", "46")
	aoc.Test(run, "input.txt", "825516882", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {

	pgs := aoc.InputParagraphs()
	//seeds := util.NumberFields(pgs[0][0])
	seeds := strings.Fields(pgs[0][0])[1:]
	pgs = pgs[1:]

	fmt.Println("seeds:", seeds)

	maps := make([][]func(string) (string, bool), 0)
	for _, pg := range pgs {
		//ff := strings.Fields(pg[0])
		//mapName := ff[0]
		mapsInner := make([]func(string) (string, bool), 0)
		for i := 1; i < len(pg); i++ {
			fff := util.NumberFields(pg[i])
			srcRangeStart := fff[1]
			destRangeStart := fff[0]
			rangeLength := fff[2]
			mapsInner = append(mapsInner, func(s string) (string, bool) {
				v := util.UnsafeAtoi(s)
				if v >= srcRangeStart && v <= srcRangeStart+rangeLength {
					return strconv.Itoa(destRangeStart + (v - srcRangeStart)), true
				} else {
					return s, false
				}
			})
		}
		maps = append(maps, mapsInner)
	}

	var found bool
	var lowest uint64 = math.MaxInt64
	for _, seed := range seeds {
		v := seed
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
		lowest = min(lowest, uint64(util.UnsafeAtoi(v)))
		fmt.Println("seed", seed, v)
	}

	//iterc.SlidingWindow(iterc.ListIterator(seeds), 2).ForEach(func(w []string) {
	//	offset := util.UnsafeAtoi(w[0])
	//	rangelen := util.UnsafeAtoi(w[1])
	//	terminus := offset + rangelen
	//	for i := offset; i < terminus; i++ {
	//		if i%1000000 == 0 {
	//			fmt.Println("million") // 20,000 of these
	//		}
	//
	//	}
	//})

	*p1 = strconv.Itoa(int(lowest))
	*p2 = strconv.Itoa(0)
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
