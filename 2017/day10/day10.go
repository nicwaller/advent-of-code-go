package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"encoding/hex"
	"strconv"
)

func main() {
	aoc.Select(2017, 10)
	aoc.Test(run, "sample.txt", "12", "")
	aoc.Test(run, "input.txt", "1980", "")
	aoc.TestLiteral(run, "AoC 2017", "", "33efeb34ea91902bb2f59c9920caa6cd")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	seq := util.NumberFields(aoc.InputString())
	isSample := len(seq) == 4
	ropeLength := map[bool]int{false: 256, true: 5}[isSample]
	rope := iter.Range(0, ropeLength).List()
	pos := 0
	skip := 0
	for _, l := range seq {
		hashround(rope, l, &pos, &skip)
	}
	*p1 = strconv.Itoa(rope[0] * rope[1])

	if isSample {
		return
	}

	rope = iter.Range(0, ropeLength).List()
	pos = 0
	skip = 0
	input := aoc.InputString()
	runes := make([]rune, len(input))
	for i, r := range input {
		runes[i] = r
	}
	runes = append(runes, 17, 31, 73, 47, 23)
	for i := 0; i < 64; i++ {
		for _, r := range runes {
			hashround(rope, int(r), &pos, &skip)
		}
	}
	*p2 = densify(rope)
}

func densify(rope []int) string {
	denseHash := make([]byte, 16)
	chunks := iter.Chunk(16, iter.ListIterator(rope))
	i := 0
	for chunks.Next() {
		c := chunks.Value()
		xorResult := f8l.Reduce(c, 0, func(a int, b int) int { return a ^ b })
		denseHash[i] = byte(xorResult)
		i++
	}
	return hex.EncodeToString(denseHash)
}

func hashround(rope []int, l int, pos *int, skip *int) {
	ropeLength := len(rope)
	rev := make([]int, ropeLength)

	// reverse copy into rev
	for j := 0; j < l; j++ {
		wIdx := j % ropeLength
		rIdx := (ropeLength + *pos + l - j - 1) % ropeLength
		rev[wIdx] = rope[rIdx]
	}
	// then copy it back into the rope
	for j := 0; j < l; j++ {
		rope[(*pos+j)%ropeLength] = rev[j%ropeLength]
	}
	*pos = (*pos + l + *skip) % ropeLength
	*skip++
}
