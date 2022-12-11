package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2016, 7)
	aoc.Test(run, "sample.txt", "2", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func hasABBA(s string) bool {
	slide := iter.SlidingWindow(4, iter.StringIterator(s, 1))
	for slide.Next() {
		ss := slide.Value()
		if ss[0] != ss[1] && ss[0] == ss[3] && ss[1] == ss[2] {
			return true
		}
	}
	return false
}

func supportIPv7(s string) bool {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '[' || r == ']'
	})
	segments := map[bool][]string{
		false: make([]string, 0),
		true:  make([]string, 0),
	}
	hyper := false
	for _, p := range parts {
		segments[hyper] = append(segments[hyper], p)
		hyper = !hyper
	}
	a := iter.
		Map[string, bool](iter.ListIterator(segments[false]), hasABBA).
		Reduce(func(a bool, b bool) bool { return a || b }, false)
	b := iter.
		Map[string, bool](iter.ListIterator(segments[true]), hasABBA).
		Reduce(func(a bool, b bool) bool { return a || b }, false)
	return a && !b
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(aoc.InputLinesIterator().Filter(supportIPv7).Count())
	*p2 = strconv.Itoa(0)
}
