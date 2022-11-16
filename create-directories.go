package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for year := 2015; year <= 2022; year++ {
		_ = os.Mkdir(strconv.Itoa(year), 0750)
		for day := 1; day <= 31; day++ {
			folderName := fmt.Sprintf("%d/day%02d", year, day)
			_ = os.Mkdir(folderName, 0750)
			safelyCreate(folderName + "/input.txt")
			safelyCreate(folderName + "/sample.txt")
			safelyCreate(folderName + "/part1.md")
			safelyCreate(folderName + "/part2.md")
			safelyCreate(folderName + "/" + folderName + ".go")
		}
	}
}

func safelyCreate(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Creating " + filename)
		_, _ = os.Create(filename)
	}
}
