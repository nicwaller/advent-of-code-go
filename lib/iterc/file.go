package iterc

import (
	"bufio"
	"os"
)

func ReadLines(filename string) (Iterator[string], error) {
	elements := make(chan string)

	file, err := os.Open(filename)
	if err != nil {
		return EmptyIterator[string](), err
	}
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			elements <- scanner.Text()
		}
		file.Close()
		close(elements)
	}()
	return Iterator[string]{
		C: elements,
	}, nil
}

func MustReadLines(filename string) Iterator[string] {
	iter, err := ReadLines(filename)
	if err != nil {
		panic(err)
	}
	return iter
}

func ReadParagraphs(filename string) (Iterator[[]string], error) {
	elements := make(chan []string)

	lines, err := ReadLines(filename)
	if err != nil {
		return EmptyIterator[[]string](), err
	}

	go func() {
		p := make([]string, 0)
		for line := range lines.C {
			if line != "" {
				p = append(p, line)
			} else {
				if len(p) > 0 {
					elements <- p
					p = make([]string, 0)
				}
			}
		}
		if len(p) > 0 {
			elements <- p
		}
		close(elements)
	}()

	return Iterator[[]string]{
		C: elements,
	}, nil
}

func MustReadParagraphs(filename string) Iterator[[]string] {
	iter, err := ReadParagraphs(filename)
	if err != nil {
		panic(err)
	}
	return iter
}
