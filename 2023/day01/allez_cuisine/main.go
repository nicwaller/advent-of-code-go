package main

import "fmt"

// Don't waste a variable on something trivial like a file handle.
//go:generate ./generate.sh

const (
	counter = 0x000000000000FFFF // 16 bits (technically only 15) for the counter
	sum     = 0x00000000FFFF0000 // 16 bits for the rolling sum
	first   = 0x0000000F00000000 // 4 bits for the first value in a line
	last    = 0x000000F000000000 // 4 bits for the last value in a line
	//temp    = 0x00FFFF0000000000 // 16 bits high for temp
)
const (
	sumShift   = 16
	firstShift = 32
	lastShift  = 36
	//tempShift  = 40
)

func main() {
	fmt.Printf("Part 1: %d\n", score(false))
	fmt.Printf("Part 2: %d\n", score(true))
}

func score(fancy bool) int {
	// The rules of "Allez Cuisine" say I can only use two variables.
	// Why is modern web-dev so wasteful? I only need one.
	var x int

	for x = 0; x&counter < len(input); x++ {
		x = (x | first) ^ first                              // zero it out
		x |= valueAtPosition(x&counter, fancy) << firstShift // save found value
		if x&first>>firstShift == 0xF {                      // nothing? try again
			continue
		}
		// we got the first value!

		for ; input[x&counter] != '\n'; x++ {
			// jump to end of line
		}

		// walk backward to find the last value on the line
		for ; ; x-- {
			x = (x | last) ^ last                               // zero it out
			x |= valueAtPosition(x&counter, fancy) << lastShift // save found value
			if x&last>>lastShift == errNoneFound {              // nothing? try again
				continue
			}
			// we got the last value!

			// add the calibration value to our rolling sum
			x = (x | sum) ^ sum | // zero the sum bits
				(0+
					(x&sum>>sumShift)+ // read current sum
					10*(x&first>>firstShift)+ // add first digit
					(x&last>>lastShift)+ // add last digit
					0)<<sumShift // save into sum

			for ; input[x&counter] != '\n'; x++ {
				// jump to end of line
			}
			break
		}
	}

	return x & sum >> sumShift
}

const errNoneFound = 0xF

func valueAtPosition(pos int, fancy bool) int {
	if input[pos] > '0' && input[pos] <= '9' {
		return int(input[pos] - '0')
	}

	if pos > len(input)-5 {
		return errNoneFound
	}

	if !fancy {
		return errNoneFound
	}

	switch {
	case input[pos:pos+3] == "one":
		return 1
	case input[pos:pos+3] == "two":
		return 2
	case input[pos:pos+5] == "three":
		return 3
	case input[pos:pos+4] == "four":
		return 4
	case input[pos:pos+4] == "five":
		return 5
	case input[pos:pos+3] == "six":
		return 6
	case input[pos:pos+5] == "seven":
		return 7
	case input[pos:pos+5] == "eight":
		return 8
	case input[pos:pos+4] == "nine":
		return 9
	default:
		return errNoneFound
	}
}
