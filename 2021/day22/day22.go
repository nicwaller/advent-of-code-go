package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var summary strings.Builder
	summary.WriteString("------------------------------------------------\n")
	part1 := ""
	part2 := ""
	aoc.UseSampleFile()
	fmt.Println("Running with sample")
	run(parseFile(), &part1, &part2)
	assert.Equal(part1, "590784")
	summary.WriteString("✅ Passed tests for sample\n")
	//assert.EqualAny(basinMultiplyResult, []int{1134, 1023660}, "basinMultiplyResult")

	part1 = ""
	part2 = ""
	aoc.UseRealInput()
	fmt.Println("Running with real input")
	run(parseFile(), &part1, &part2)
	//assert.EqualAny(basinMultiplyResult, []int{1134, 1023660}, "basinMultiplyResult")
	summary.WriteString(fmt.Sprintf("Part 1: %s\n", part1))
	summary.WriteString(fmt.Sprintf("Part 2: %s\n", part2))
	summary.WriteString("https://adventofcode.com/2021/day/22")
	fmt.Printf("%s", summary.String())
}

type fileType struct {
	steps []rebootStep
}

type rebootStep struct {
	state bool
	over  grid.Slice
}

func parseFile() fileType {
	f := fileType{steps: make([]rebootStep, 0)}
	for lines := aoc.InputLinesIterator(); lines.Next(); {
		line := lines.Value()
		parts := strings.Split(line, " ")
		state := parts[0] == "on"
		nParts := util.NumberFields(parts[1])
		f.steps = append(f.steps, rebootStep{
			state: state,
			over: []grid.Range{
				grid.Range{
					Origin:   nParts[0],
					Terminus: nParts[1] + 1,
				},
				grid.Range{
					Origin:   nParts[2],
					Terminus: nParts[3] + 1,
				},
				grid.Range{
					Origin:   nParts[4],
					Terminus: nParts[5] + 1,
				},
			},
		})
	}
	return f
}

func run(input fileType, p1 *string, p2 *string) {
	universe := grid.Slice{
		grid.Range{
			Origin:   -50,
			Terminus: 51,
		},
		grid.Range{
			Origin:   -50,
			Terminus: 51,
		},
		grid.Range{
			Origin:   -50,
			Terminus: 51,
		},
	}
	// TODO: newGrid() should support non-zero origin?
	g := grid.NewGridFromSlice[bool](universe)
	for _, x := range input.steps {
		fmt.Println(x)
		iSlice, err := x.over.Intersect(universe)
		if err != nil {
			fmt.Println(err)
			continue
		}
		g.FillSlice(x.state, iSlice)
	}
	cubesOn := g.Filter(func(v bool) bool { return v }).Count()
	*p1 = strconv.Itoa(cubesOn)
	return
}
