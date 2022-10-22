package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read file
	fbytes, _ := os.ReadFile("input.txt")
	fmt.Printf("Part 1: %d", part1(fbytes))
	fmt.Printf("Part 2: %d", part2(fbytes))
}

func part1(filebytes []byte) int {
	fstr := string(filebytes)
	// split by line
	lines := strings.Split(fstr, "\n")
	// strconv integer
	values := stringArrayToIntArray(lines)
	// sliding window / pairwise
	var dives int
	pairwise[int](values, func(t1 int, t2 int) {
		if t2 > t1 {
			dives++
		}
	})
	// difference
	// filter to positive values
	// count
	return dives
}

func part2(filebytes []byte) int {
	fstr := string(filebytes)
	lines := strings.Split(fstr, "\n")
	values := stringArrayToIntArray(lines)
	var dives int
	var left, right []int
	slidingWindow[int](values, 3, func(slice []int) {
		if left == nil {
			left = slice
			return
		} else {
			right = slice
			leftsum := left[0] + left[1] + left[2]
			rightsum := right[0] + right[1] + right[2]
			if rightsum > leftsum {
				dives++
			}
			left = right
			right = nil
		}
	})
	return dives
}

func stringArrayToIntArray(lines []string) []int {
	result := make([]int, len(lines))
	for index, item := range lines {
		result[index], _ = strconv.Atoi(item)
	}
	return result
}

func pairwise[T comparable](list []T, handler func(T, T)) {
	slidingWindow(list, 2, func(slice []T) {
		handler(slice[0], slice[1])
	})
}

func slidingWindow[T comparable](list []T, windowSize int, handler func([]T)) {
	for i := 0; i <= len(list)-windowSize; i++ {
		handler(list[i : i+windowSize])
	}
}
