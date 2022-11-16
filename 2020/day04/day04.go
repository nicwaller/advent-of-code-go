package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2020, 4)
	aoc.Test(run, "sample.txt", "2", "")
	aoc.Test(run, "input.txt", "254", "184")
	aoc.Run(run)
}

type passport map[string]string

func run(p1 *string, p2 *string) {
	valid1 := 0
	valid2 := 0
	iter.Map(aoc.ParagraphsIterator(), decodePassport).
		Filter(hasAllFields).Echo().Counting(&valid1).
		Filter(isValid).Counting(&valid2).Go()
	*p1 = strconv.Itoa(valid1)
	*p2 = strconv.Itoa(valid2)
}

func decodePassport(paragraph []string) passport {
	p := make(map[string]string)
	for _, part := range strings.Fields(strings.Join(paragraph, " ")) {
		k, v := util.Pair(strings.Split(part, ":"))
		p[k] = v
	}
	return p
}

func hasAllFields(p passport) bool {
	keys := set.New[string]()
	keys.Insert("cid")
	for k, _ := range p {
		keys.Insert(k)
	}
	return keys.Size() == 8
}

// isValid assumes all fields are populated
func isValid(p passport) bool {
	v := true
	v = v && util.InRange(util.UnsafeAtoi(p["byr"]), 1920, 2002)
	v = v && util.InRange(util.UnsafeAtoi(p["iyr"]), 2010, 2020)
	v = v && util.InRange(util.UnsafeAtoi(p["eyr"]), 2020, 2030)
	height := util.NumberFields(p["hgt"])[0]
	if strings.HasSuffix(p["hgt"], "cm") {
		v = v && util.InRange(height, 150, 193)
	} else if strings.HasSuffix(p["hgt"], "in") {
		v = v && util.InRange(height, 59, 76)
	} else {
		return false
	}
	v = v && util.Must(regexp.MatchString(`^#[0-9a-f]{6}$`, p["hcl"]))
	v = v && util.Must(regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, p["ecl"]))
	v = v && util.Must(regexp.MatchString(`^[0-9]{9}$`, p["pid"]))
	return v
}
