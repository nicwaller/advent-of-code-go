package main

import (
	"fmt"
	"os"
)

func main() {
	for day := 1; day <= 31; day++ {
		folderName := fmt.Sprintf("day%02d", day)
		os.Mkdir(folderName, 0750)
		safelyCreate(folderName + "/input.txt")
		safelyCreate(folderName + "/part1.md")
		safelyCreate(folderName + "/part2.md")
		safelyCreate(folderName + "/" + folderName + ".go")
	}
}

func safelyCreate(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Creating " + filename)
		os.Create(filename)
	}
}
