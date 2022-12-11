package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"strings"
)

func main() {
	aoc.Select(2016, 6)
	aoc.Test(run, "sample.txt", "easter", "advent")
	//aoc.Test(run, "input.txt", "zcreqgiv", "pljvorrk")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	g := grid.FromString(aoc.InputString())
	var sb strings.Builder
	var sb2 strings.Builder
	for c := 0; c < g.ColumnCount(); c++ {
		sb.WriteString(analyze.Analyze(g.Column(c)).MostCommon)
		sb2.WriteString(analyze.Analyze(g.Column(c)).LeastCommon)
	}
	*p1 = sb.String()
	*p2 = sb2.String()
}
