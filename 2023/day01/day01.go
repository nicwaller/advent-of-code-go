package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 1)
	aoc.Test(run, "sample.txt", "142", "")
	aoc.Test(run, "sample2.txt", "", "281")
	aoc.Test(run, "input.txt", "55090", "54845")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	sum1 := 0
	for _, line := range aoc.InputLines() {
		nf := util.Digits(line)
		if len(nf) == 0 {
			continue
		}
		nn := (nf[0]%10)*10 + (nf[len(nf)-1])%10
		sum1 += nn
	}
	*p1 = strconv.Itoa(sum1)

	sum2 := 0
	for _, line := range aoc.InputLines() {
		dd := dumbDigits(line)
		//fmt.Println(dd) // TODO: remove

		first := dd[0]
		last := dd[len(dd)-1]

		n := util.UnsafeAtoi(first)*10 + util.UnsafeAtoi(last)
		sum2 += n
	}
	*p2 = strconv.Itoa(sum2)
}

func dumbDigits(line string) []string {
	digits := make([]string, 0)
	for i, c := range line {
		if c >= '0' && c <= '9' {
			digits = append(digits, string(c))
			continue
		}

		for name, dd := range base10Digits {
			if strings.HasPrefix(line[i:], name) {
				digits = append(digits, dd)
			}
		}
	}
	return digits
}

var base10Digits = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
