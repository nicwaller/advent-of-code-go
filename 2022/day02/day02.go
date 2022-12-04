package main

import (
	"advent-of-code/lib/aoc"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 2)
	aoc.Test(run, "sample.txt", "15", "12")
	aoc.Test(run, "input.txt", "11150", "")
	aoc.Run(run)
}

const (
	OutcomeLoss int = 0
	OutcomeDraw     = 1
	OutcomeWin      = 2
)

const (
	ShapeRock     int = 1
	ShapePaper        = 2
	ShapeScissors     = 3
)

func score(myShape int, outcome int) int {
	return myShape + outcome*3
}

func outcome(opponentShape int, myShape int) int {
	lookup := opponentShape<<2 | myShape
	ret := map[int]int{
		ShapeRock<<2 | ShapeRock:         OutcomeDraw,
		ShapeRock<<2 | ShapePaper:        OutcomeWin,
		ShapeRock<<2 | ShapeScissors:     OutcomeLoss,
		ShapePaper<<2 | ShapeRock:        OutcomeLoss,
		ShapePaper<<2 | ShapePaper:       OutcomeDraw,
		ShapePaper<<2 | ShapeScissors:    OutcomeWin,
		ShapeScissors<<2 | ShapeRock:     OutcomeWin,
		ShapeScissors<<2 | ShapePaper:    OutcomeLoss,
		ShapeScissors<<2 | ShapeScissors: OutcomeDraw,
	}[lookup]
	return ret
}

func resolveShape(s string) int {
	return map[string]int{
		"A": ShapeRock,
		"B": ShapePaper,
		"C": ShapeScissors,
		"X": ShapeRock,
		"Y": ShapePaper,
		"Z": ShapeScissors,
	}[s]
}

func playForOutcome(oppoShape int, desiredOutcome int) int {
	// I know this can be done cleverly with math and maybe modulus,
	// but I am sick, so I am doing this.
	for myPlay := 1; myPlay <= 3; myPlay++ {
		if outcome(oppoShape, myPlay) == desiredOutcome {
			return myPlay
		}
	}
	panic("The only way to win is not to play")
}

func run(p1 *string, p2 *string) {
	myScore := 0
	myScore2 := 0
	for _, line := range aoc.InputLines() {
		f := strings.Fields(line)
		oppo := resolveShape(f[0])
		mine := resolveShape(f[1])
		o := outcome(oppo, mine)
		s := score(mine, o)
		myScore += s

		desiredOutcome := map[string]int{
			"X": OutcomeLoss,
			"Y": OutcomeDraw,
			"Z": OutcomeWin,
		}[f[1]]
		myPlay2 := playForOutcome(oppo, desiredOutcome)
		myScore2 += score(myPlay2, desiredOutcome)
	}
	*p1 = strconv.Itoa(myScore)
	*p2 = strconv.Itoa(myScore2)
}
