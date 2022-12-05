package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/stack"
	"advent-of-code/lib/util"
	"fmt"
	"strings"
)

func main() {
	aoc.Select(2022, 5)
	//aoc.Test(run, "sample.txt", "CMZ", "")
	aoc.Test(run, "input.txt", "TGWSMRBPN", "")
	aoc.Run(run)
	aoc.Out()
}

func parseStates(lines []string) []stack.Stack[string] {
	all := strings.Join(lines, "\n")
	g := grid.FromString(all)

	stacks := make([]stack.Stack[string], 9)
	for j := 0; j < len(stacks); j++ {
		stacks[j] = stack.NewStack[string]()
	}
	for x := 1; x < 34; x += 4 {
		xID := x / 4
		for y := 7; y >= 0; y-- {
			v := g.Get(grid.Cell{y, x})
			if v != " " {
				stacks[xID].Push(v)
			}
		}
	}

	return stacks
}

func run(p1 *string, p2 *string) {
	state, moves := util.Pair(aoc.ParagraphsIterator().List())
	s := parseStates(state)
	s2 := parseStates(state)
	_ = moves
	fmt.Println(s)

	for _, mv := range moves {
		nf := util.NumberFields(mv)
		cnt := nf[0]
		src := nf[1]
		dst := nf[2]
		for i := 0; i < cnt; i++ {
			v := s[src-1].MustPop()
			s[dst-1].Push(v)
		}
		vList := s2[src-1].MustPopN(cnt)
		for i := len(vList) - 1; i >= 0; i-- {
			s2[dst-1].Push(vList[i])
		}

	}

	var sb strings.Builder
	for _, sss := range s {
		sb.WriteString(sss.Peek())
	}

	var sb2 strings.Builder
	for _, sss := range s2 {
		sb2.WriteString(sss.Peek())
	}

	*p1 = sb.String()
	*p2 = sb2.String()
}
