package aoc

import (
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"fmt"
	"os"
	"strings"
	"time"
)

var inputFilename = "input.txt"

type RunFunc func(p1 *string, p2 *string)

var summary strings.Builder
var dayNumber int

func Select(year int, day int) {
	dayNumber = day
	wd, _ := os.Getwd()
	dayDir := fmt.Sprintf("%d/day%02d", year, day)
	if wd == dayDir {
		// good!
	} else {
		err := os.Chdir(dayDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func Out() {
	if summary.Len() == 0 {
		return
	}
	fmt.Println("---------------------------------------------------")
	fmt.Print(summary.String())
	fmt.Printf("Submit: https://adventofcode.com/2021/day/%d", dayNumber)
	summary.Reset()
}

//func IsSample() bool {
//	return inputFilename != "input.txt"
//}

func Run(run RunFunc) {
	// TODO: timings!
	var p1Actual string
	var p2Actual string
	inputFilename = "input.txt"
	start := time.Now()
	run(&p1Actual, &p2Actual)
	elapsed := time.Since(start).Round(time.Millisecond)
	summary.WriteString(fmt.Sprintf("Completed in %v\n", elapsed))
	summary.WriteString(fmt.Sprintf("Part 1: %v\n", p1Actual))
	summary.WriteString(fmt.Sprintf("Part 2: %v\n", p2Actual))
	Out()
}

func Test(run RunFunc, filename string, p1 string, p2 string) {
	inputFilename = filename
	_, err := os.Stat(filename)
	if err != nil {
		summary.WriteString("❌ Missing File (" + filename + ")\n")
		return
	}
	var p1Actual string
	var p2Actual string
	run(&p1Actual, &p2Actual)
	if p1 == "" {
		// skip
	} else if p1 == p1Actual {
		summary.WriteString(fmt.Sprintf(
			"%s %-12s part %d (\"%s\")\n",
			"✅  ok  ", filename, 1, p1Actual))

	} else {
		summary.WriteString(fmt.Sprintf(
			"%s %-12s part %d (expected %s but got %s)\n",
			"❌  fail", filename, 1, p1, p1Actual))
	}
	if p2 == "" {
		//summary.WriteString("⏸ Ignoring result from part 2\n")
	} else if p2 == p2Actual {
		summary.WriteString(fmt.Sprintf(
			"%s %-12s part %d (\"%s\")\n",
			"✅  ok  ", filename, 2, p2Actual))
	} else {
		summary.WriteString(fmt.Sprintf(
			"%s %-12s part %d (expected %s but got %s)\n",
			"❌  fail", filename, 2, p2, p2Actual))
	}
}

func UseFile(filename string) {
	inputFilename = filename
}
func UseSampleFile() {
	inputFilename = "sample.txt"
}
func UseRealFile() {
	inputFilename = "input.txt"
}

//func InputFilename() string {
//	return inputFilename
//}

//goland:noinspection GoUnusedExportedFunction
func InputBytes() []byte {
	file, err := os.ReadFile(inputFilename)
	if err != nil {
		panic(err)
	}
	if len(file) == 0 {
		panic("empty input")
	}
	return file
}

//goland:noinspection GoUnusedExportedFunction
func InputString() string {
	return strings.TrimSpace(string(InputBytes()))
}

//goland:noinspection GoUnusedExportedFunction
func InputLines() []string {
	return util.ReadLines(inputFilename).List()
}

//goland:noinspection GoUnusedExportedFunction
func InputLinesIterator() iter.Iterator[string] {
	return util.ReadLines(inputFilename)
}

//goland:noinspection GoUnusedExportedFunction
func InputGridRunes() grid.Grid[string] {
	s := InputString()
	return grid.FromString(s)
}

//goland:noinspection GoUnusedExportedFunction
func InputGridNumbers() grid.Grid[int] {
	s := InputString()
	return grid.FromStringAsInt(s)
}
