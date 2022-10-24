package util

import (
	"bufio"
	"bytes"
	"os"
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

func ReadLines(filename string) []string {
	fbytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(fbytes))
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ScanMulti(scanner *bufio.Scanner, into *[]string) bool {
	for i := 0; i < len(*into); i++ {
		if !scanner.Scan() {
			return false
		}
		(*into)[i] = scanner.Text()
	}
	return true
}
