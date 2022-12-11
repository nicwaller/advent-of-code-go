package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2016, 5)
	//aoc.Test(run, "sample.txt", "18f47a30", "05ace8e3")
	//aoc.Test(run, "input.txt", "801b56a7", "424a0197")
	aoc.Run(run)
	aoc.Out()
}

func interestingHashes(prefix string) iter.Iterator[string] {
	c := -1
	var md5hex string
	return iter.Iterator[string]{
		Next: func() bool {
			for {
				c++
				plaintext := fmt.Sprintf("%s%d", prefix, c)
				md5bytes := md5.Sum([]byte(plaintext))
				md5hex = hex.EncodeToString(md5bytes[:])
				if strings.HasPrefix(md5hex, "00000") {
					return true
				}
			}
		},
		Value: func() string {
			return md5hex
		},
	}
}

func run(p1 *string, p2 *string) {
	*p1 = strings.Join(interestingHashes(aoc.InputString()).
		Take(8).
		Map(func(s string) string { return s[5:6] }).
		List(), "")

	password := make([]string, 8)
	for h := interestingHashes(aoc.InputString()); h.Next(); {
		s := h.Value()
		ix, err := strconv.Atoi(s[5:6])
		if err != nil {
			continue
		}
		if ix > 7 {
			continue
		}
		if password[ix] == "" {
			password[ix] = s[6:7]
		}
		fmt.Println(password)
		if slices.Index(password, "") == -1 {
			*p2 = strings.Join(password, "")
			break
		}
	}
}
