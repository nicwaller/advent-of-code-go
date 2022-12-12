package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2016, 14)
	aoc.TestLiteral(run, "abc", "22728", "22551")
	//aoc.Test(run, "input.txt", "23890", "22696")
	aoc.Run(run)
}

func hashes(prefix string, stretch int) iter.Iterator[string] {
	index := 0 // I expected this to be -1. Why does it need 0?
	var md5hex string
	return iter.Iterator[string]{
		Next: func() bool {
			for {
				index++
				plaintext := fmt.Sprintf("%s%d", prefix, index)
				md5bytes := md5.Sum([]byte(plaintext))
				md5hex = hex.EncodeToString(md5bytes[:])
				for i := 0; i < stretch; i++ {
					reBytes := md5.Sum([]byte(md5hex))
					md5hex = hex.EncodeToString(reBytes[:])
				}
				return true
			}
		},
		Value: func() string {
			return md5hex
		},
	}
}

func interestingCandidates(s string) []string {
	candidates := make([]string, 0)
	window := iter.SlidingWindow(3, iter.StringIterator(s, 1))
	for window.Next() {
		v := window.Value()
		if v[0] == v[1] && v[1] == v[2] {
			candidates = append(candidates, v[0])
			// only consider the FIRST such triplet of a hash
			return candidates
		}
	}
	return candidates
}

func run(p1 *string, p2 *string) {
	salt := aoc.InputString()
	keys := make([]string, 0)
	window := iter.Enumerate[[]string](iter.SlidingWindow(1000, hashes(salt, 0)))
scanLoop1:
	for window.Next() {
		index, v := window.Value().Pair()
		maybeKey := v[0]
		candidates := interestingCandidates(maybeKey)
		for _, followHash := range v[1:] {
			for _, c := range candidates {
				if strings.Contains(followHash, strings.Repeat(c, 5)) {
					keys = append(keys, maybeKey)
					continue scanLoop1
				}
			}
		}
		if len(keys) == 64 {
			*p1 = strconv.Itoa(index)
			break
		}
	}

	keys = make([]string, 0)
	window = iter.Enumerate[[]string](iter.SlidingWindow(1000, hashes(salt, 2016)))
scanLoop2:
	for window.Next() {
		index, v := window.Value().Pair()
		maybeKey := v[0]
		candidates := interestingCandidates(maybeKey)
		for _, followHash := range v[1:] {
			for _, c := range candidates {
				if strings.Contains(followHash, strings.Repeat(c, 5)) {
					keys = append(keys, maybeKey)
					continue scanLoop2
				}
			}
		}
		if len(keys) == 64 {
			*p2 = strconv.Itoa(index)
			break
		}
	}
}
