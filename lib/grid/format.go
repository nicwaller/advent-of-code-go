package grid

import (
	"fmt"
	"strings"
)

func (grid *Grid[T]) Print() {
	fmt.Println(grid.String())
}

func (grid *Grid[T]) String() string {
	switch len(grid.dimensions) {
	case 1:
		return fmt.Sprint(grid.storage)
	case 2:
		rowLength := grid.dimensions[1]
		rows := len(grid.storage) / rowLength
		var rowStrings []string
		for row := 0; row < rows; row++ {
			offset := row * rowLength
			line := fmt.Sprint(grid.storage[offset : offset+rowLength])
			rowStrings = append(rowStrings, line)
		}
		return strings.Join(rowStrings, "\n")
	default:
		return fmt.Sprintf("Cannot print %d-dimensional grid", len(grid.dimensions))
	}
	// be mindful of dimensions
	// maybe redo this with reader/writer
}
