package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2015, 8)
	aoc.Test(run, "input.txt", "1342", "2074")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	rawSize := iter.Map[string, int](aoc.InputLinesIterator(), func(s string) int { return len(s) }).
		Reduce(util.IntSum, 0)
	finalSize := iter.Map[string, int](aoc.InputLinesIterator().Map(unescape), func(s string) int { return len(s) }).
		Reduce(util.IntSum, 0)
	overhead := rawSize - finalSize
	*p1 = strconv.Itoa(overhead)

	finalSize2 := iter.Map[string, int](aoc.InputLinesIterator().Map(escape), func(s string) int { return len(s) }).
		Reduce(util.IntSum, 0)
	overhead2 := finalSize2 - rawSize
	*p2 = strconv.Itoa(overhead2)
}

const (
	StateRegular          = 0
	StateSequenceOpen     = 1
	StateSequenceLiteral  = 2
	StateSequenceLiteral2 = 2
)

func unescape(s string) string {
	var sb strings.Builder
	state := StateRegular
	const backslash = '\\'
	var first rune
	for _, r := range s[1 : len(s)-1] {
		switch state {
		case StateRegular:
			switch {
			case r == backslash:
				state = StateSequenceOpen
				continue
			default:
				sb.WriteRune(rune(r))
			}
		case StateSequenceOpen:
			switch {
			case r == backslash:
				fallthrough
			case r == '"':
				sb.WriteRune(rune(r))
				state = StateRegular
				continue
			case r == 'x':
				state = StateSequenceLiteral
				first = 0
			default:
				fmt.Println(s)
				fmt.Println(first, r)
				panic("invalid escape sequence")
			}
		case StateSequenceLiteral:
			switch {
			case r >= '0' && r <= '9':
				fallthrough
			case r >= 'a' && r <= 'f':
				if first == 0 {
					first = rune(r)
				} else {
					hexstr := string(first) + string(r)
					b, _ := hex.DecodeString(hexstr)
					sb.Write(b)
					state = StateRegular
				}
			default:
				panic("invalid sequence literal")
			}
		}
	}
	z := sb.String()
	fmt.Println(z)
	return z
}

func escape(s string) string {
	var sb strings.Builder
	sb.WriteString(`"`)
	for _, r := range s {
		switch r {
		case '\\':
			sb.WriteString(`\\`)
		case '"':
			sb.WriteString(`\"`)
		default:
			sb.WriteRune(r)
		}
	}
	sb.WriteString(`"`)
	return sb.String()
}
