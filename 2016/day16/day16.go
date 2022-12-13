package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"strings"
)

var disklen int

func main() {
	aoc.Select(2016, 16)
	disklen = 20
	aoc.Test(run, "sample.txt", "01100", "10111110011110111")
	disklen = 272
	aoc.Test(run, "input.txt", "10100011010101011", "")
	aoc.Run(run)
}

func parse() []uint8 {
	m := map[rune]uint8{
		'0': 0,
		'1': 1,
	}
	s := aoc.InputString()
	ret := make([]uint8, len(s))
	for i, r := range s {
		ret[i] = m[r]
	}
	return ret
}

func expand(s []uint8) []uint8 {
	ret := make([]uint8, len(s)*2+1)
	copy(ret, s)
	o := len(ret) - 1
	for _, v := range s {
		ret[o] = v ^ 1
		o--
	}
	return ret
}

func checksum(s []uint8) []uint8 {
	ret := make([]uint8, len(s)/2)
	chunks := iter.Chunk(2, iter.ListIterator(s))
	for i := 0; chunks.Next(); i++ {
		c := chunks.Value()
		ret[i] = c[0] ^ c[1] ^ 1
	}
	if len(ret)%2 == 0 {
		return checksum(ret)
	} else {
		return ret
	}
}

func format(s []uint8) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for _, v := range s {
		sb.WriteByte('0' + v)
	}
	return sb.String()
}

func run(p1 *string, p2 *string) {
	dat := parse()
	for len(dat) < disklen {
		dat = expand(dat)
	}
	dat = dat[0:disklen]
	*p1 = format(checksum(dat))

	dat = parse()
	for len(dat) < 35651584 {
		dat = expand(dat)
	}
	dat = dat[0:35651584]
	*p2 = format(checksum(dat))
}
