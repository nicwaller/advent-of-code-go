package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2016, 4)
	aoc.Test(run, "sample.txt", "1514", "")
	aoc.Test(run, "input.txt", "158835", "")
	aoc.Run(run)
}

type F struct {
	char        rune
	occurrences int
}

func checksum(s string) string {
	counts := util.Make[F](26, func(i int) F {
		return F{
			char:        rune('a' + i),
			occurrences: 0,
		}
	})
	for _, c := range strings.Replace(s, "-", "", -1) {
		i := c - 'a'
		counts[i].occurrences++
	}
	sort.SliceStable(counts, func(i, j int) bool {
		return counts[i].occurrences > counts[j].occurrences
	})
	res := strings.Join(f8l.Map[F, string](counts[0:5], func(f F) string { return string(f.char) }), "")
	return res
}

func decrypt(name string, sector int) string {
	var sb strings.Builder
	for _, c := range name {
		if c == '-' {
			sb.WriteRune('-')
			continue
		}
		alpha := int(c - 'a')
		alpha = (alpha + sector) % 26
		sb.WriteRune(rune(alpha + 'a'))
	}
	return sb.String()
}

func run(p1 *string, p2 *string) {
	r := regexp.MustCompile(`^([a-z-]*)([0-9]*)\[(.....)\]$`)

	sum := 0
	for _, line := range aoc.InputLines() {
		m := r.FindStringSubmatch(line)
		roomNameCiphertext := m[1]
		if checksum(roomNameCiphertext) == m[3] {
			sector := util.UnsafeAtoi(m[2])
			sum += sector
			room := decrypt(roomNameCiphertext, sector)
			if strings.Contains(room, "north") {
				fmt.Println(room)
				*p2 = strconv.Itoa(sector)
			}
		}
	}
	*p1 = strconv.Itoa(sum)
}
