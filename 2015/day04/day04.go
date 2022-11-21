package main

import (
	"advent-of-code/lib/aoc"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2015, 4)
	aoc.Test(run, "sample.txt", "609043", "")
	aoc.Test(run, "input.txt", "254575", "1038736")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	input := aoc.InputString()
	for i := 1; ; i++ {
		s := fmt.Sprintf("%s%d", input, i)
		b := md5.Sum([]byte(s))
		if strings.HasPrefix(hex.EncodeToString(b[:]), "00000") {
			*p1 = strconv.Itoa(i)
			break
		}
	}
	for i := 1; ; i++ {
		s := fmt.Sprintf("%s%d", input, i)
		b := md5.Sum([]byte(s))
		if strings.HasPrefix(hex.EncodeToString(b[:]), "000000") {
			*p2 = strconv.Itoa(i)
			break
		}
	}
}
