package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/queue"
	"advent-of-code/lib/stack"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

var isSample bool

func main() {
	aoc.Select(2016, 10)
	isSample = true
	aoc.Test(run, "sample.txt", "2", "")
	isSample = false
	//aoc.Test(run, "input.txt", "73", "3965")
	aoc.Run(run)
	aoc.Out()
}

type Bot struct {
	inventory         queue.Queue[int]
	lowTarget         int
	highTarget        int
	lowTargetIsOutput bool
}

func parse() []Bot {
	botRules := aoc.InputLinesIterator().Filter(func(s string) bool { return strings.HasPrefix(s, "bot") }).List()
	bots := make([]Bot, len(botRules))
	for _, rule := range botRules {
		nf := util.NumberFields(rule)
		botId := nf[0]
		bots[botId] = Bot{
			inventory:         queue.New[int](2),
			lowTarget:         nf[1],
			highTarget:        nf[2],
			lowTargetIsOutput: strings.Fields(rule)[5] == "output",
		}
	}
	seedValues := aoc.InputLinesIterator().Filter(func(s string) bool { return strings.HasPrefix(s, "value") })
	for seedValues.Next() {
		nf := util.NumberFields(seedValues.Value())
		bots[nf[1]].inventory.Push(nf[0])
	}
	return bots
}

func run(p1 *string, p2 *string) {
	outputs := util.Make(21, func(i int) stack.Stack[int] {
		return stack.NewStack[int]()
	})
	bots := parse()
	keepGoing := true
	for keepGoing {
		keepGoing = false
		for botID, _ := range bots {
			bot := &bots[botID]
			if bot.inventory.Length() == 2 {
				keepGoing = true
				a := analyze.Analyze(bot.inventory.Items())
				if isSample && a.Min == 2 && a.Max == 5 {
					*p1 = strconv.Itoa(botID)
				}
				if !isSample && (a.Min == 17 && a.Max == 61) {
					*p1 = strconv.Itoa(botID)
				}
				bot.inventory.Reset()
				if bot.lowTargetIsOutput {
					outputs[bot.lowTarget].Push(a.Min)
				} else {
					bots[bot.lowTarget].inventory.Push(a.Min)
				}
				bots[bot.highTarget].inventory.Push(a.Max)
			}
		}
	}
	if !isSample {
		*p2 = strconv.Itoa(util.IntProductV(
			outputs[0].MustPop(),
			outputs[1].MustPop(),
			outputs[2].MustPop(),
		))
	}
}
