package aoc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"
	"time"

	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/iterc"
	"advent-of-code/lib/util"
)

var inputFilename = "input.txt"

type RunFunc func(p1 *string, p2 *string)

var summary strings.Builder
var dayNumber int
var yearNumber int
var allTestsPassed = true

func Select(year int, day int) {
	dayNumber = day
	yearNumber = year
	wd, _ := os.Getwd()
	dayDir := fmt.Sprintf("%d/day%02d", year, day)
	if wd == dayDir {
		// good!
	} else {
		err := os.Chdir(dayDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	//MaybeDownload()
}

func MaybeDownload() {
	getCookie := func() string {
		homeDir := util.Must(user.Current()).HomeDir
		cookieFile := path.Join(homeDir, ".aoc", "session")
		fileBytes := util.Must(os.ReadFile(cookieFile))
		return strings.TrimSpace(string(fileBytes))
	}
	hasContent := func(filename string) bool {
		stat, err := os.Stat(filename)
		return err == nil && stat.Size() > 0
	}
	download := func(url string, toFile string) {
		dstFile := util.Must(os.Create(toFile))
		req := util.Must(http.NewRequest("GET", url, nil))
		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: getCookie(),
		})
		httpResp := util.Must(http.DefaultClient.Do(req))
		if httpResp.StatusCode != 200 {
			fmt.Println(httpResp)
			fmt.Println(url)
			fmt.Println(httpResp.Status)
			os.Exit(1)
		}
		util.Must(io.Copy(dstFile, httpResp.Body))
	}
	if !hasContent("input.txt") {
		url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", yearNumber, dayNumber)
		fmt.Println(url)
		download(fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", yearNumber, dayNumber), "input.txt")
		fmt.Println("Downloaded input.txt")
	}
}

func Out() {
	if summary.Len() == 0 {
		return
	}
	fmt.Println("---------------------------------------------------")
	fmt.Print(summary.String())
	if allTestsPassed {
		fmt.Printf("Submit: https://adventofcode.com/%d/day/%d \n", yearNumber, dayNumber)
	}
	summary.Reset()
}

//func IsSample() bool {
//	return inputFilename != "input.txt"
//}

func Run(run RunFunc) {
	// TODO: timings!
	var p1Actual string
	var p2Actual string
	inputFilename = "input.txt"
	if !allTestsPassed {
		summary.WriteString(fmt.Sprintf("⏸️Won't run %s until tests are passing", inputFilename))
		Out()
		return
	}
	start := time.Now()
	fmt.Printf("🏃Running: %s\n", inputFilename)
	run(&p1Actual, &p2Actual)
	elapsed := time.Since(start).Round(time.Millisecond)
	summary.WriteString(fmt.Sprintf("Completed in %v\n", elapsed))
	summary.WriteString(fmt.Sprintf("Part 1: %v\n", p1Actual))
	summary.WriteString(fmt.Sprintf("Part 2: %v\n", p2Actual))
	Out()
}

func TestLiteral(run RunFunc, content string, p1 string, p2 string) {
	file, err := os.CreateTemp("", "aoc")
	if err != nil {
		log.Fatal(err)
	}
	_ = util.Must(file.WriteString(content))
	defer func(name string) {
		_ = os.Remove(name)
	}(file.Name())
	Test(run, file.Name(), p1, p2)
}

func Test(run RunFunc, filename string, p1 string, p2 string) {
	inputFilename = filename
	st, err := os.Stat(filename)
	if err != nil {
		summary.WriteString("❌ Missing File (" + filename + ")\n")
		allTestsPassed = false
		return
	}
	if st.Size() == 0 {
		summary.WriteString("❌ Empty File (" + filename + ")\n")
		allTestsPassed = false
		return
	}
	fmt.Printf("🔍Testing: %s\n", filename)
	var p1Actual string
	var p2Actual string
	run(&p1Actual, &p2Actual)
	if p1 == "" {
		// skip
	} else if p1 == p1Actual {
		summary.WriteString(fmt.Sprintf(
			"%s %-12s part %d (%q)\n",
			"✅  ok  ", filename, 1, p1Actual))

	} else {
		allTestsPassed = false
		summary.WriteString(fmt.Sprintf(
			"%s %-12s part %d (expected %q but got %q)\n",
			"❌  fail", filename, 1, p1, p1Actual))
	}
	if p2 == "" {
		//summary.WriteString("⏸ Ignoring result from part 2\n")
	} else if p2 == p2Actual {
		summary.WriteString(fmt.Sprintf(
			"%s %-12s part %d (%q)\n",
			"✅  ok  ", filename, 2, p2Actual))
	} else {
		allTestsPassed = false
		summary.WriteString(fmt.Sprintf(
			"%s %-12s part %d (expected %q but got %q)\n",
			"❌  fail", filename, 2, p2, p2Actual))
	}
}

func UseFile(filename string) {
	inputFilename = filename
}
func UseSampleFile() {
	inputFilename = "sample.txt"
}
func UseRealFile() {
	inputFilename = "input.txt"
}

//func InputFilename() string {
//	return inputFilename
//}

//goland:noinspection GoUnusedExportedFunction
func InputBytes() []byte {
	file, err := os.ReadFile(inputFilename)
	if err != nil {
		panic(err)
	}
	if len(file) == 0 {
		panic("empty input")
	}
	return file
}

//goland:noinspection GoUnusedExportedFunction
func InputString() string {
	return strings.TrimSpace(string(InputBytes()))
}

//goland:noinspection GoUnusedExportedFunction
func InputLines() []string {
	return iterc.MustReadLines(inputFilename).List()
}

//goland:noinspection GoUnusedExportedFunction
func InputLinesInt() []int {
	return iterc.Map(iterc.MustReadLines(inputFilename), util.UnsafeAtoi).List()
}

//goland:noinspection GoUnusedExportedFunction
func InputLinesIterator() iter.Iterator[string] {
	return util.ReadLines(inputFilename)
}

//goland:noinspection GoUnusedExportedFunction
func InputLinesIterc() iterc.Iterator[string] {
	return iterc.MustReadLines(inputFilename)
}

func InputParagraphs() [][]string {
	return iterc.MustReadParagraphs(inputFilename).List()
}

//goland:noinspection GoUnusedExportedFunction
func InputGridRunes() grid.Grid[string] {
	s := InputString()
	return grid.FromString(s)
}

//goland:noinspection GoUnusedExportedFunction
func InputGridNumbers() grid.Grid[int] {
	s := InputString()
	return grid.FromStringAsInt(s)
}

func ParagraphsIterator() iter.Iterator[[]string] {
	lines := InputLinesIterator()
	var parLines []string
	return iter.Iterator[[]string]{
		Next: func() bool {
			parLines = lines.TakeWhile(func(v string) bool { return v != "" }).List()
			return len(parLines) > 0
		},
		Value: func() []string {
			return parLines
		},
	}
}

func InputDelimited2(d1, d2 string) [][]string {
	// ignore whitespace
	unwrap := func(a, b string) string { return a + b }
	input := InputLinesIterator().Reduce(unwrap, "")

	results := make([][]string, 0)

	parts1 := strings.Split(input, d1)
	for _, part := range parts1 {
		results = append(results, strings.Split(part, d2))
	}
	return results
}
