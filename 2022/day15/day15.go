package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"fmt"
	"math"
	"strconv"
)

var checkRow int
var searchSpace int

func main() {
	aoc.Select(2022, 15)
	checkRow = 10
	searchSpace = 20
	aoc.Test(run, "sample.txt", "26", "56000011")
	checkRow = 2000000
	searchSpace = 4000000
	aoc.Test(run, "input.txt", "5166077", "")
	//aoc.Run(run)
	aoc.Out()
}

type SegmentedRange[T comparable] struct {
	segments []grid.Range
	context  []T
}

func (sr *SegmentedRange[T]) Init() {
	var none T
	sr.context = []T{none}
	sr.segments = []grid.Range{
		{Origin: math.MinInt32, Terminus: math.MaxInt32},
	}
}

func (sr *SegmentedRange[T]) At(x int) (int, T) {
	// PERF: binary search would be faster!
	for i, seg := range sr.segments {
		if x > seg.Origin {
			return i, sr.context[i]
		}
	}
	panic("no matching segment")
}

func (sr *SegmentedRange[T]) SetAll(r grid.Range, v T) {
	// PERF: binary search would still be faster...
	for i, seg := range sr.segments {
		if seg.Origin >= r.Origin && seg.Terminus <= r.Terminus {
			sr.context[i] = v
		}
		if seg.Origin >= r.Terminus {
			return
		}
	}
}

func (sr *SegmentedRange[T]) Cut(x int) {
	sCopy := make([]grid.Range, 0)
	cCopy := make([]T, 0)

	for i, seg := range sr.segments {
		if seg.Terminus < x {
			// pre-segment
			sCopy = append(sCopy, sr.segments[i])
			cCopy = append(cCopy, sr.context[i])
		} else if seg.Origin > x {
			// post-segment
			sCopy = append(sCopy, sr.segments[i])
			cCopy = append(cCopy, sr.context[i])
		} else {
			if x == seg.Terminus {
				//fmt.Printf("Skipping cut at %d\n", x)
				return
			}
			// overlapping segment
			lHalf := grid.Range{
				Origin:   seg.Origin,
				Terminus: x,
			}
			rHalf := grid.Range{
				Origin:   x,
				Terminus: seg.Terminus,
			}
			sCopy = append(sCopy, lHalf, rHalf)
			cCopy = append(cCopy, sr.context[i], sr.context[i])
		}
	}
	sr.segments = sCopy
	sr.context = cCopy
}

func (sr *SegmentedRange[T]) Count(v T) int {
	sum := 0
	for i, s := range sr.segments {
		if s.Terminus-s.Origin > 10000000 {
			// HAX
			continue
		}
		if sr.context[i] == v {
			sum += s.Terminus - s.Origin
		}
	}
	return sum
}

type Reading struct {
	sensor grid.Cell
	beacon grid.Cell
}

func row(y int, readings []Reading) SegmentedRange[bool] {
	sr := SegmentedRange[bool]{}
	sr.Init()
	for _, r := range readings {
		d0 := grid.ManhattanDistance(r.sensor, r.beacon)
		yOffset := util.IntAbs(r.sensor[0] - y)
		d1 := d0 - yOffset
		if d1 < 0 {
			continue
		}
		x0 := r.sensor[1] - d1
		x1 := r.sensor[1] + d1
		sr.Cut(x0)
		sr.Cut(x1)
		sr.SetAll(grid.Range{
			Origin:   x0,
			Terminus: x1,
		}, true)
	}
	return sr
}

func run(p1 *string, p2 *string) {
	sr := SegmentedRange[rune]{}
	sr.Init()

	readings := make([]Reading, 0)
	for _, line := range aoc.InputLines() {
		nf := util.NumberFields(line)
		readings = append(readings, Reading{
			sensor: grid.Cell{nf[1], nf[0]},
			beacon: grid.Cell{nf[3], nf[2]},
		})
	}
	r := row(checkRow, readings)
	c := r.Count(true)
	*p1 = strconv.Itoa(c)

	//for y := 0; y <= searchSpace; y++ {
	for y := 2703981; y <= 2703981; y++ {
		rr := row(y, readings)
		v := rr.Count(false)
		_ = rr
		_ = v
		if v > 0 {
			fmt.Printf("v=%d y=%d %v\n", v, y, rr.context)
			for i := 1; i < len(rr.context)-1; i++ {
				if rr.context[i] == false {
					fmt.Println(rr.segments[i-1])
					fmt.Println(rr.segments[i])
					fmt.Println(rr.segments[i+1])
					tuningFrequency := 4000000*rr.segments[i].Origin + y
					*p2 = strconv.Itoa(tuningFrequency)
				}
			}
		}
	}
}
