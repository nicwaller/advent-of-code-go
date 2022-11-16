package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"strconv"
)

func main() {
	aoc.Select(2021, 3)
	aoc.Test(run, "sample.txt", "198", "230")
	aoc.Test(run, "input.txt", "4138664", "4273224")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(part1())
	*p2 = strconv.Itoa(part2())
}

func gammaRate(g grid.Grid[int]) int {
	// Each bit in the gamma rate can be determined by
	// finding the most common bit in the corresponding position
	// of all numbers in the diagnostic report.
	gammaRate := 0
	bitPos := g.ColumnCount() - 1
	for column := g.ColumnIter(); column.Next(); bitPos-- {
		gammaRate |= analyze.MostCommon(column.Value()) << bitPos
	}
	return gammaRate
}

func epsilonRate(gammaRate int, bits int) int {
	// The epsilon rate is calculated in a similar way;
	// rather than use the most common bit,
	// the least common bit from each position is used.
	epsilonMask := 0
	for bit := 0; bit < bits; bit++ {
		epsilonMask += 1 << bit
	}
	epsilonRate := gammaRate ^ epsilonMask
	return epsilonRate
}

func part1() int {
	g := grid.FromStringAsInt(aoc.InputString())
	gammaRate := gammaRate(g)
	epsilonRate := epsilonRate(gammaRate, g.ColumnCount())
	return gammaRate * epsilonRate
}

func p2_oxygenGeneratorRating(g grid.Grid[int]) int {
	// Create a Set() of numbers for easy pruning later
	decodeBinaryDigits := func(row []int) int {
		sum := 0
		for i := 0; i < len(row); i++ {
			sum += row[len(row)-i-1] << i
		}
		return sum
	}
	numbers := set.FromIterable(iter.Map(g.RowIter(), decodeBinaryDigits))

	// Start calculating oxygen generator rating
	// Keep only numbers with most common value in each bit Position
	bitSize := g.ColumnCount()
	for bitPos := 0; bitPos < bitSize; bitPos++ {
		bitShift := bitSize - bitPos - 1
		bitMask := 1 << bitShift
		sum := 0

		//getBit := func(v int) int {
		//	return binary.NthBitI(bitPos, v)
		//}
		//desBit := analyze.MostCommon(numbers.Iterator().Map(getBit).List())

		for _, item := range numbers.Items() {
			if item&bitMask > 0 {
				sum += 1
			}
		}
		var desiredBitStatus int
		if sum*2 >= numbers.Size() {
			desiredBitStatus = 1
		} else {
			desiredBitStatus = 0
		}
		//assert.Equal(desBit, desiredBitStatus)
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
	}
	assert.Equal(numbers.Size(), 1)
	return numbers.Items()[0]
}

func part2() int {
	g := grid.FromStringAsInt(aoc.InputString())

	bitSize := g.ColumnCount()

	// Set up a Set() of numbers for easy pruning later
	decodeBinaryDigits := func(row []int) int {
		sum := 0
		for i := 0; i < len(row); i++ {
			sum += row[len(row)-i-1] << i
		}
		return sum
	}
	numbers := set.FromIterable(iter.Map(g.RowIter(), decodeBinaryDigits))
	numbersCopy := set.Union(numbers)

	oxygenGeneratorRating := p2_oxygenGeneratorRating(g)
	//assert.EqualNamed(oxygenGeneratorRating, 2031, "oxygenGeneratorRating")

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

	co2ScrubberRating := numbers.Items()[0]
	lifeSupportRating := oxygenGeneratorRating * co2ScrubberRating
	return lifeSupportRating
}
