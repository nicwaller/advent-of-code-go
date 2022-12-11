package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
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

func supernetHypernet(s string) ([]string, []string) {
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
	return segments[false], segments[true]
}

func supportIPv7(s string) bool {
	supernets, hypernets := supernetHypernet(s)
	a := iter.
		Map[string, bool](iter.ListIterator(supernets), hasABBA).
		Reduce(func(a bool, b bool) bool { return a || b }, false)
	b := iter.
		Map[string, bool](iter.ListIterator(hypernets), hasABBA).
		Reduce(func(a bool, b bool) bool { return a || b }, false)
	return a && !b
}

func supportsSSL(s string) bool {
	abaSet := set.New[string]()
	sup, hyp := supernetHypernet(s)
	for _, s := range sup {
		slide := iter.SlidingWindow(3, iter.StringIterator(s, 1))
		for slide.Next() {
			v := slide.Value()
			if v[0] == v[2] && v[0] != v[1] {
				abaSet.Add(strings.Join(v, ""))
			}
		}
	}
	for _, aba := range abaSet.Items() {
		bab := aba[1:2] + aba[0:1] + aba[1:2]
		for _, h := range hyp {
			if strings.Contains(h, bab) {
				return true
			}
		}
	}
	return false
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(aoc.InputLinesIterator().Filter(supportIPv7).Count())
	*p2 = strconv.Itoa(aoc.InputLinesIterator().Filter(supportsSSL).Count())
}
