package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/stack"
	"fmt"
	"strconv"
)

func main() {
	aoc.Select(2017, 9)
	aoc.TestLiteral(run, "{}", "1", "")
	aoc.TestLiteral(run, "{{{}}}", "6", "")
	aoc.TestLiteral(run, "{{},{}}", "5", "")
	aoc.TestLiteral(run, "{<a>,<a>,<a>,<a>}", "1", "")
	aoc.TestLiteral(run, "{{<ab>},{<ab>},{<ab>},{<ab>}}", "9", "")
	aoc.TestLiteral(run, "<>", "", "0")
	aoc.TestLiteral(run, "<random characters>", "", "17")
	aoc.TestLiteral(run, "<<<<>", "", "3")
	aoc.TestLiteral(run, "<{!>}>", "", "2")
	aoc.TestLiteral(run, "<!!>", "", "0")
	aoc.TestLiteral(run, "<!!!>>", "", "0")
	aoc.TestLiteral(run, "<{o\"i!a,<{i<a>", "", "10")
	//aoc.Test(run, "input.txt", "17390", "7825")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	score := 0
	s := stack.NewStack[string]()
	stream := iter.StringIterator(aoc.InputString(), 1)
	garbage := 0
	var isGarbage bool
	for stream.Next() {
		n := stream.Value()
		fmt.Printf("%s", n)
		if n == "!" {
			stream.Next()
			fmt.Printf("\033[31;7m")
			fmt.Printf(stream.Value())
			fmt.Printf("\u001B[0m")
			continue
		}
		if isGarbage && n != ">" {
			garbage++
			continue
		}
		switch n {
		case "{":
			s.Push("{")
		case "}":
			score += s.Height()
			s.MustPop()
		case "<":
			isGarbage = true
			fmt.Printf("\033[31;1;4m")
		case ">":
			isGarbage = false
			fmt.Printf("\u001B[0m")
		case ",":
			// more things!
		default:
			if isGarbage {
				// oh well
			} else {
				fmt.Printf("\n")
				//fmt.Printf("what the fuck is a %q?\n", n)
				panic(n)
			}
		}
	}
	*p1 = strconv.Itoa(score)
	*p2 = strconv.Itoa(garbage)
	fmt.Printf("\n")
}
