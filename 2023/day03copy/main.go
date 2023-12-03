package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type engine [][]rune

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	eng := make(engine, 0, len(lines))
	for _, line := range lines {
		eng = append(eng, []rune(line))
	}

	seen := map[[2]int]bool{}

	part1 := 0
	part2 := 0

	for y := 0; y < len(eng); y++ {
		for x := 0; x < len(eng[y]); x++ {
			isNumeric := eng[y][x] >= '0' && eng[y][x] <= '9'
			if !isNumeric && eng[y][x] != '.' {
				isGear := eng[y][x] == '*'
				nextTo := []int{}
				for yd := -1; yd <= 1; yd++ {
					for xd := -1; xd <= 1; xd++ {
						x := x + xd
						y := y + yd
						if yd == 0 && xd == 0 || y < 0 || y >= len(eng) || x < 0 || x >= len(eng[y]) {
							continue
						}

						isNumeric := eng[y][x] >= '0' && eng[y][x] <= '9'
						if isNumeric {
							startX := x
							for ; startX >= 0 && eng[y][startX] >= '0' && eng[y][startX] <= '9'; startX-- {
							}
							startX += 1

							endX := x
							for ; endX < len(eng[y]) && eng[y][endX] >= '0' && eng[y][endX] <= '9'; endX++ {
							}

							if seen[[2]int{y, startX}] {
								continue
							}

							number, err := strconv.Atoi(string(eng[y][startX:endX]))
							if err != nil {
								panic(err)
							}
							fmt.Println(number)

							part1 += number
							nextTo = append(nextTo, number)
							seen[[2]int{y, startX}] = true
						}
					}
				}

				if len(nextTo) == 2 && isGear {
					part2 += nextTo[0] * nextTo[1]
				}
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
