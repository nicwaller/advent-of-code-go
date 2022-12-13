package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func main() {
	aoc.Select(2016, 17)
	aoc.TestLiteral(run, "ihgpwlah", "DDRRRD", "370")
	aoc.TestLiteral(run, "kglvqrro", "DDUDRLRRUDRD", "492")
	aoc.TestLiteral(run, "ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR", "830")
	//aoc.Test(run, "input.txt", "DUDRLRRDDR", "")
	aoc.Run(run)
	aoc.Out()
}

func isOpen(r uint8) bool {
	return strings.ContainsAny(string(r), "bcdef")
}

var dirNames = []string{"U", "D", "L", "R"}
var dirVecs = map[string]grid.Cell{
	"U": {-1, 0},
	"D": {1, 0},
	"L": {0, -1},
	"R": {0, 1},
}

func getPath(ttl int, prefix string, pos grid.Cell, goal grid.Cell, path string) string {
	if ttl == 0 {
		return ""
	}
	if pos[0] == goal[0] && pos[1] == goal[1] {
		return path
	}
	plaintext := prefix + path
	md5sum := md5.Sum([]byte(plaintext))
	md5hex := hex.EncodeToString(md5sum[:])[0:4]
	for i, dir := range dirNames {
		if isOpen(md5hex[i]) {
			newPos := make([]int, 2)
			copy(newPos, pos)
			util.VecAdd(newPos, dirVecs[dir])
			if newPos[0] < 0 || newPos[0] >= 4 || newPos[1] < 0 || newPos[1] >= 4 {
				continue
			}
			deep := getPath(ttl-1, prefix, newPos, goal, path+dir)
			if deep != "" {
				return deep
			}
		}
	}
	return ""
}

func run(p1 *string, p2 *string) {
	input := aoc.InputString()
	for i := 1; ; i++ {
		r := getPath(i, input, grid.Cell{0, 0}, grid.Cell{3, 3}, "")
		if r != "" && *p1 == "" {
			*p1 = r
			break
		}
	}
}
