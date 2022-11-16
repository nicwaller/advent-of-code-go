package util

import (
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"bufio"
	"golang.org/x/exp/constraints"
	"math"
	"os"
	"regexp"
	"strconv"
)

func UnsafeAtoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func Take[T any](count int, from []T) ([]T, []T) {
	offset := count
	take := from[0:offset]
	remainder := from[offset:]
	return take, remainder
}

func TakeOne[T any](from []T) (T, []T) {
	first, remaining := Take[T](1, from)
	return first[0], remaining
}

func Chunk[T any](list []T, chunkSize int) [][]T {
	numChunks := len(list) / chunkSize
	chunks := make([][]T, numChunks)
	offset := 0
	for i := 0; i < numChunks; i++ {
		chunks[i] = list[offset : offset+chunkSize]
		offset += chunkSize
	}
	return chunks
}

//func ReadLines(filename string) []string {
//	fbytes, err := os.ReadFile(filename)
//	if err != nil {
//		panic(err)
//	}
//	scanner := bufio.NewScanner(bytes.NewReader(fbytes))
//	lines := make([]string, 0)
//	for scanner.Scan() {
//		lines = append(lines, scanner.Text())
//	}
//	return lines
//}

func ScanMulti(scanner *bufio.Scanner, into *[]string) bool {
	for i := 0; i < len(*into); i++ {
		if !scanner.Scan() {
			return false
		}
		(*into)[i] = scanner.Text()
	}
	return true
}

func ReadLines(filename string) iter.Iterator[string] {
	file, err := os.Open(filename)
	if err != nil {
		// could also return EmptyIterator?
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var line string
	return iter.Iterator[string]{
		Next: func() bool {
			if !scanner.Scan() {
				_ = file.Close()
				return false
			}
			line = scanner.Text()
			return true
		},
		Value: func() string {
			return line
		},
	}
}

// NumberFields is like strings.Fields() but it gets all the integers
// including negative integers ;)
// and be wary of stray hyphens in the input!
// TODO: rename to IntFields()
func NumberFields(s string) []int {
	var intMatcher = regexp.MustCompile("-?[0-9]+")
	matches := intMatcher.FindAllString(s, math.MaxInt32)
	return f8l.Map[string, int](matches, UnsafeAtoi)
	//return
	//stringFields := strings.FieldsFunc(s, func(r rune) bool {
	//	return r != '-' && (r < '0' || r > '9')
	//})
	//intFields := make([]int, len(stringFields))
	//for i := 0; i < len(intFields); i++ {
	//	intFields[i], _ = strconv.Atoi(stringFields[i])
	//}
	//return intFields
}

func IntMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func IntMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//goland:noinspection GoUnusedExportedFunction
func IntSum(a int, b int) int {
	return a + b
}

//goland:noinspection GoUnusedExportedFunction
func IntProduct(a int, b int) int {
	return a * b
}

//goland:noinspection GoUnusedExportedFunction
func IntProductV(a ...int) int {
	r := 1
	for _, v := range a {
		r *= v
	}
	return r
}

//goland:noinspection GoUnusedExportedFunction
func IntAbs(a int) int {
	if a < 0 {
		return 0 - a
	} else {
		return a
	}
}

func IntIncr(a int) int {
	return a + 1
}

func Identity[T any](v T) T {
	return v
}

func GreaterThan[T constraints.Ordered](a T, b T) int {
	if a > b {
		return 1
	}
	return 0
}

func LessThan[T constraints.Ordered](a T, b T) int {
	if a < b {
		return 1
	}
	return 0
}

// FromHexChar converts a hex character into its value and a success flag.
func FromHexChar(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	panic(c)
}

//goland:noinspection GoUnusedExportedFunction
func Must[T any](result T, err9 error) T {
	if err9 != nil {
		panic(err9)
	}
	return result
}

func Pair[T any](tList []T) (T, T) {
	if len(tList) != 2 {
		panic(len(tList))
	}
	return tList[0], tList[1]
}
