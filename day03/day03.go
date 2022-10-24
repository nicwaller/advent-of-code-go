package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/set"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content := parseFile()
	fmt.Printf("Part 1: %d\n", part1(content))
	fmt.Printf("Part 2: %d\n", part2(content))
}

func parseFile() grid.Grid[int] {
	fbytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	return grid.FromStringAsInt(string(fbytes))
}

func gammaRate(g grid.Grid[int]) int {
	// Each bit in the gamma rate can be determined by
	// finding the most common bit in the corresponding position
	// of all numbers in the diagnostic report.
	var mostPrevalentBits = make([]string, g.ColumnCount())
	for c := 0; c < g.ColumnCount(); c++ {
		column := g.Column(c)
		mostPrevalentBits[c] = strconv.Itoa(analyze.MostCommon(column))
	}
	binStr := strings.Join(mostPrevalentBits, "")
	gammaRate, _ := strconv.ParseInt(binStr, 2, 32)
	return int(gammaRate) // 22 for sample, ?? for input
}

func epsilonRate(gammaRate int, bits int) int {
	// The epsilon rate is calculated in a similar way;
	// rather than use the most common bit,
	// the least common bit from each position is used.
	epsilonMask := 0
	for bit := 0; bit < bits; bit++ {
		epsilonMask += 1 << bit
	}
	return gammaRate ^ epsilonMask
}

func part1(g grid.Grid[int]) int {
	gammaRate := gammaRate(g)
	epsilonRate := epsilonRate(gammaRate, g.ColumnCount())
	// The power consumption can then be found by multiplying the gamma rate by the epsilon rate.
	powerConsumption := gammaRate * epsilonRate
	fmt.Printf("gammaRate: %d\n", gammaRate)
	fmt.Printf("epsilonRate: %d\n", epsilonRate)
	fmt.Printf("powerConsumption: %d\n", powerConsumption)
	return powerConsumption
}

func part2(g grid.Grid[int]) int {
	bitSize := g.ColumnCount()

	// Set up a Set() of numbers for easy pruning later
	var numbers = set.New[int]()
	nextRow := g.RowIterator()
	for row := nextRow(); row != nil; row = nextRow() {
		binDigits := fmap[int, string](row, func(i int) string {
			return strconv.Itoa(i)
		})
		binStr := strings.Join(binDigits, "")
		binVal, _ := strconv.ParseInt(binStr, 2, 32)
		numbers.Add(int(binVal))
	}

	numbersCopy := set.Union(numbers)

	// Start calculating oxygen generator rating
	// Keep only numbers with most common value in each bit Position
	for bitPos := 0; bitPos < bitSize; bitPos++ {
		//fmt.Println("-----------------------")
		//fmt.Printf("nth bit: %d\n", bitPos)
		bitShift := bitSize - bitPos - 1
		//fmt.Printf("bit shift: %d\n", bitShift)
		bitMask := 1 << bitShift
		//fmt.Printf("bitMask: %s\n", strconv.FormatInt(int64(bitMask), 2))
		sum := 0
		for _, item := range numbers.Items() {
			if item&bitMask > 0 {
				sum += 1
			}
		}
		//fmt.Printf("In pos %d, found %d/%d bits are positive\n", bitPos, sum, numbers.Size())
		var desiredBitStatus int
		if sum*2 >= numbers.Size() {
			desiredBitStatus = 1
		} else {
			desiredBitStatus = 0
		}
		//fmt.Printf("desiredBitStatus: %d\n", desiredBitStatus)
		numbers.Filter(func(v int) bool {
			bit := v & bitMask
			switch desiredBitStatus {
			case 0:
				return bit == 0
			case 1:
				return bit > 0
			default:
				panic(1)
			}
		})
		//fmt.Println(numbers.String())
	}
	if numbers.Size() != 1 {
		panic("should have been only 1 number left")
	}
	oxygenGeneratorRating := numbers.Items()[0]

	// Restore the original number set
	numbers = numbersCopy

	// Start calculating CO2 Scrubber Rating
	// Keep only numbers with most common value in each bit Position
	for bitPos := 0; bitPos < bitSize; bitPos++ {
		//fmt.Println("-----------------------")
		//fmt.Printf("nth bit: %d\n", bitPos)
		bitShift := bitSize - bitPos - 1
		//fmt.Printf("bit shift: %d\n", bitShift)
		bitMask := 1 << bitShift
		//fmt.Printf("bitMask: %s\n", strconv.FormatInt(int64(bitMask), 2))
		sum := 0
		for _, item := range numbers.Items() {
			if item&bitMask > 0 {
				sum += 1
			}
		}
		//fmt.Printf("In pos %d, found %d/%d bits are positive\n", bitPos, sum, numbers.Size())
		var desiredBitStatus int
		if sum*2 >= numbers.Size() {
			desiredBitStatus = 0
		} else {
			desiredBitStatus = 1
		}
		//fmt.Printf("desiredBitStatus: %d\n", desiredBitStatus)
		numbers.Filter(func(v int) bool {
			bit := v & bitMask
			switch desiredBitStatus {
			case 0:
				return bit == 0
			case 1:
				return bit > 0
			default:
				panic(1)
			}
		})
		//fmt.Println(numbers.String())
		if numbers.Size() == 1 {
			break
		}
	}
	if numbers.Size() != 1 {
		panic("should have been only 1 number left")
	}
	co2ScrubberRating := numbers.Items()[0]

	lifeSupportRating := oxygenGeneratorRating * co2ScrubberRating

	fmt.Printf("oxygenGeneratorRating: %d\n", oxygenGeneratorRating)
	fmt.Printf("co2ScrubberRating: %d\n", co2ScrubberRating)
	fmt.Printf("lifeSupportRating: %d\n", lifeSupportRating)

	return lifeSupportRating // 4273224 for the input
}

func fmap[I interface{}, O interface{}](items []I, mapFn func(I) O) []O {
	var results []O = make([]O, len(items))
	for _, item := range items {
		results = append(results, mapFn(item))
	}
	return results
}
