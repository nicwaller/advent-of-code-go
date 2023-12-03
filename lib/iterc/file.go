package iterc

import (
	"bufio"
	"os"
)

func MustReadLines(filename string) Iterator[string] {
	iter, err := ReadLines(filename)
	if err != nil {
		panic(err)
	}
	return iter
}

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
