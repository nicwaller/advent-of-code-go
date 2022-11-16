package main

import (
	"advent-of-code/lib/util"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

func main() {
	baseDir, _ := os.Getwd()
	for year := 2015; year <= 2022; year++ {
		yearDir := filepath.Join(baseDir, strconv.Itoa(year))
		_ = os.Mkdir(yearDir, 0750)
		for day := 1; day <= 31; day++ {
			dayDirName := fmt.Sprintf("day%02d", day)
			dayDir := filepath.Join(yearDir, dayDirName)
			_ = os.Mkdir(dayDir, 0750)
			if err := os.Chdir(dayDir); err != nil {
				panic(err)
			}
			safelyCreate("input.txt")
			safelyCreate("sample.txt")
			safelyCreate("part1.md")
			safelyCreate("part2.md")
			safelyCreateContent(dayDirName+".go", skel(year, day))
		}
	}
}

func safelyCreate(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Creating " + filename)
		util.Must(os.Create(filename))
	}
}

func safelyCreateContent(filename string, content []byte) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Creating %s (%d bytes)\n", filename, len(content))
		if err := os.WriteFile(filename, content, 0644); err != nil {
			panic(err)
		}
	}
}

func skel(year int, day int) []byte {
	gotmpl := `package main

import (
	"advent-of-code/lib/aoc"
	"strconv"
)

func main() {
	aoc.Select({{.Year}}, {{.Day}})
	aoc.Test(run, "sample.txt", "", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(0)
	*p2 = strconv.Itoa(0)
}
`
	tmpl, err := template.New("skel").Parse(gotmpl)
	if err != nil {
		panic(err)
	}
	type when struct {
		Year int
		Day  int
	}
	buff := new(bytes.Buffer)
	err = tmpl.Execute(buff, when{
		Year: year,
		Day:  day,
	})
	return buff.Bytes()
}
