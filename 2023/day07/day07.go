package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iterc"
	"advent-of-code/lib/util"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2023, 7)
	aoc.Test(run, "sample.txt", "6440", "5905")
	aoc.Test(run, "input.txt", "249483956", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	plays := make([]Play, 0)
	for _, line := range aoc.InputLines() {
		cards, bidStr := util.Pair(strings.Fields(line))
		plays = append(plays, Play{
			cards: []rune(cards),
			hand:  BestHand([]rune(cards)),
			bid:   util.UnsafeAtoi(bidStr),
		})
	}

	//fmt.Println(plays)
	slices.SortFunc(plays, func(a, b Play) bool {
		if a.hand != b.hand {
			return a.hand < b.hand
		}
		for i := range a.cards {
			if d := CardValues[a.cards[i]] - CardValues[b.cards[i]]; d != 0 {
				return d < 0
			}
		}
		return false
		//if analyze.Analyze(a.cards).Max < analyze.Analyze(b.cards).Max {
		//	return true
		//}
		//if a.cards[0] < b.cards[0] {
		//	return true
		//}
		//return false
	})
	//fmt.Println(plays)

	winnings := 0
	iterc.EnumerateFrom(iterc.ListIterator(plays), 1).ForEach(func(iv iterc.IndexedValue[Play]) {
		rank := iv.Index
		p := iv.Value
		fmt.Println(rank, string(p.cards), p.bid, p.hand)
		winnings += rank * p.bid
	})

	*p1 = strconv.Itoa(winnings)
	*p2 = strconv.Itoa(0)
}

type Play struct {
	cards []rune
	bid   int
	hand  Hand
}

func BestHand(cards []rune) Hand {
	a := analyze.Analyze(cards)
	switch a.CountMostCommon {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if a.CountLeastCommon == 2 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	case 2:
		if a.Distinct == 3 {
			return TwoPair
		} else {
			return OnePair
		}
	default:
		return HighCard
	}
}

type Hand uint8

const (
	HighCard     Hand = iota
	OnePair           = iota
	TwoPair           = iota
	ThreeOfAKind      = iota
	FullHouse         = iota
	FourOfAKind       = iota
	FiveOfAKind       = iota
)

var CardValues = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}
