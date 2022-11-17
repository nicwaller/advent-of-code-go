package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"strconv"
)

func main() {
	aoc.Select(2020, 6)
	aoc.Test(run, "sample.txt", "11", "6")
	aoc.Test(run, "input.txt", "6549", "3466")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	count := 0
	for para := aoc.ParagraphsIterator(); para.Next(); {
		s := set.New[string]()
		for _, pers := range para.Value() {
			iter.StringIterator(pers, 1).Each(func(answer string) {
				s.Insert(answer)
			})
		}
		count += s.Size()
	}
	*p1 = strconv.Itoa(count)

	count2 := 0
	for para := aoc.ParagraphsIterator(); para.Next(); {
		peeps := para.Value()
		groupYesAnswers := set.FromIterable(iter.StringIterator(peeps[0], 1))
		for _, pers := range peeps[1:] {
			persAnswers := set.FromIterable(iter.StringIterator(pers, 1))
			groupYesAnswers = set.Intersection(groupYesAnswers, persAnswers)
		}
		count2 += groupYesAnswers.Size()
	}
	*p2 = strconv.Itoa(count2)
}
