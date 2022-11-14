package aoc

import (
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"os"
	"strings"
)

var filename = "input.txt"

func UseSampleFile() {
	filename = "sample.txt"
}
func UseRealInput() {
	filename = "input.txt"
}
func InputFilename() string {
	return filename
}

//goland:noinspection GoUnusedExportedFunction
func InputBytes() []byte {
	file, err := os.ReadFile(InputFilename())
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
	return util.ReadLines(InputFilename()).List()
}

//goland:noinspection GoUnusedExportedFunction
func InputLinesIterator() iter.Iterator[string] {
	return util.ReadLines(InputFilename())
}

func InputGridRunes() grid.Grid[string] {
	s := InputString()
	return grid.FromString(s)
}

func InputGridNumbers() grid.Grid[int] {
	s := InputString()
	return grid.FromStringAsInt(s)
}
