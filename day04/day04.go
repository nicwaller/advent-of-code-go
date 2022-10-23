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

type fileType struct {
	draws  []int
	boards []bingoBoard
}
type bingoBoard [][]int

func parseFile() fileType {
	//arrayify (return 2d array of string)
	//Map each item in the grid to int
	//
	//For each draw number
	// For each board
	//  grid.replace(draw, NaN) or -1
	//  isWinner() ?
	//  sumUnmarked() * drawNumber
	fbytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	//Split by line
	scanner := bufio.NewScanner(bytes.NewReader(fbytes))

	//Take 1 for Number stream, split by comma, parseInt.
	// numbers drawn
	scanner.Scan()
	drawNums := fmap[string, int](strings.Split(scanner.Text(), ","), unsafeAtoi)

	// bingo boards
	//Divide remainder into chunks of 6 lines
	var boards []bingoBoard
	for {
		boardLines := scanMulti(scanner, 6)
		if len(boardLines) == 0 {
			break
		}
		//Prune the first (empty) line
		boardLines = boardLines[1:]

		var board = make([][]int, 5)
		for rowIdx, rowStr := range boardLines {
			fields := strings.Fields(rowStr)
			vals := fmap(fields, unsafeAtoi)
			board[rowIdx] = vals
		}

		boards = append(boards, board)
	}
	return fileType{
		draws:  drawNums,
		boards: boards,
	}
}

func scanMulti(scanner *bufio.Scanner, count int) []string {
	results := make([]string, count)
	for i := 0; i < count; i++ {
		if !scanner.Scan() {
			return []string{}
		}
		results[i] = scanner.Text()
		//results = append(results, scanner.Text())
	}
	return results
}

func unsafeAtoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func checkBingo(board bingoBoard) bool {
	possibleBingo := true
	for y := 0; y < 5; y++ {
		possibleBingo = true
		for x := 0; x < 5 && possibleBingo; x++ {
			if board[y][x] != 100 {
				possibleBingo = false
				continue
			}
		}
		if possibleBingo {
			return true
		}
	}
	for x := 0; x < 5; x++ {
		possibleBingo = true
		for y := 0; y < 5 && possibleBingo; y++ {
			if board[y][x] != 100 {
				possibleBingo = false
				continue
			}
		}
		if possibleBingo {
			return true
		}
	}
	return false
}

func boardSum(board bingoBoard) int {
	sum := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			val := board[y][x]
			if val == 100 {
				continue
			} else {
				sum += val
			}
		}
	}
	return sum
}

func part1(input fileType) int {
	for _, drawn := range input.draws {
		for _, board := range input.boards {
			changed := false
			for rowId, row := range board {
				for colId, val := range row {
					if val == drawn {
						board[rowId][colId] = 100
						changed = true
					}
				}
			}
			if changed {
				if checkBingo(board) {
					finalScore := drawn * boardSum(board)
					return finalScore
				}
			}
		}
	}
	return -1
}

func part2(input fileType) int {
	var lastWinScore int
	alreadyBingo := make([]bool, len(input.boards))
	for drawNumber, drawn := range input.draws {
		for boardId, board := range input.boards {
			if alreadyBingo[boardId] {
				continue
			}
			changed := false
			for rowId, row := range board {
				for colId, val := range row {
					if val == drawn {
						board[rowId][colId] = 100
						changed = true
					}
				}
			}
			if changed {
				if checkBingo(board) {
					fmt.Printf("got a bingo! on draw %d\n", drawNumber)
					lastWinScore = drawn * boardSum(board)
					alreadyBingo[boardId] = true
				}
			}
		}
	}
	return lastWinScore // 4880
}

func fmap[I interface{}, O interface{}](items []I, mapFn func(I) O) []O {
	var results []O = make([]O, len(items))
	for i, item := range items {
		results[i] = mapFn(item)
		//results = append(results, mapFn(item))
	}
	return results
}
