package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2016, 9)
	aoc.TestLiteral(run, "ADVENT", "6", "")
	aoc.TestLiteral(run, "A(1x5)BC", "7", "")
	aoc.TestLiteral(run, "(3x3)XYZ", "9", "9")
	aoc.TestLiteral(run, "A(2x2)BCD(2x2)EFG", "11", "")
	aoc.TestLiteral(run, "(6x1)(1x3)A", "6", "")
	aoc.TestLiteral(run, "X(8x2)(3x3)ABCY", "18", "20")
	aoc.TestLiteral(run, "(27x12)(20x12)(13x14)(7x10)(1x12)A", "", "241920")
	aoc.TestLiteral(run, "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", "", "445")
	//aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
	aoc.Out()
}

func decompressV1(s string) string {
	var sb strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			c := strings.Index(s[i:], ")")
			size, repeats := util.Pair(f8l.Map(strings.Split(s[i+1:i+c], "x"), util.UnsafeAtoi))
			chunk := s[i+c+1 : i+c+1+size]
			for i := 0; i < repeats; i++ {
				sb.WriteString(chunk)
			}
			i = i + c + size
		} else {
			sb.WriteByte(s[i])
		}
		i++
	}
	return sb.String()
}

func decompressV1len(s string) int {
	slen := 0
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			c := strings.Index(s[i:], ")")
			size, repeats := util.Pair(f8l.Map(strings.Split(s[i+1:i+c], "x"), util.UnsafeAtoi))
			slen += repeats * size
			i = i + c + size
		} else {
			slen++
		}
		i++
	}
	return slen
}

func decompressV2(s string) string {
	var sb strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			c := strings.Index(s[i:], ")")
			size, repeats := util.Pair(f8l.Map(strings.Split(s[i+1:i+c], "x"), util.UnsafeAtoi))
			chunk := s[i+c+1 : i+c+1+size]
			for i := 0; i < repeats; i++ {
				sb.WriteString(chunk)
			}
			i = i + c + size
		} else {
			sb.WriteByte(s[i])
		}
		i++
	}
	r := sb.String()
	if strings.Index(r, "(") > -1 {
		return decompressV2(r)
	} else {
		return sb.String()
	}
}

func decompressV2len(s string) int {
	slen := 0
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			c := strings.Index(s[i:], ")")
			size, repeats := util.Pair(f8l.Map(strings.Split(s[i+1:i+c], "x"), util.UnsafeAtoi))
			chunk := s[i+c+1 : i+c+1+size]
			slen += repeats * decompressV2len(chunk)
			i = i + c + size
		} else {
			slen++
		}
		i++
	}
	return slen
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(len(decompressV1(aoc.InputString())))
	*p2 = strconv.Itoa(decompressV2len(aoc.InputString()))
}
