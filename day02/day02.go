package main

import (
	"bufio"
	"bytes"
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

type fileType []record
type record struct {
	direction string
	quantity  int
}

func parseFile() fileType {
	fbytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var records []record
	scanner := bufio.NewScanner(bytes.NewReader(fbytes))
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		distance, _ := strconv.Atoi(fields[1])
		records = append(records, record{
			direction: fields[0],
			quantity:  distance,
		})
	}
	return records
}

func part1(input fileType) int {
	var y, z = 0, 0
	for _, movement := range input {
		switch movement.direction {
		case "up":
			z -= movement.quantity
		case "down":
			z += movement.quantity
		case "forward":
			y += movement.quantity
		}
	}

	return y * z // 1580000
}

func part2(input fileType) int {
	var y, z = 0, 0
	var aim = 0
	for _, movement := range input {
		switch movement.direction {
		case "down":
			aim += movement.quantity
		case "up":
			aim -= movement.quantity
		case "forward":
			y += movement.quantity
			z += movement.quantity * aim
		}
	}

	return y * z
}
