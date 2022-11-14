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
		var sb strings.Builder
		sb.Grow(len(grid.storage))
		for offset, val := range grid.storage {
			sb.WriteString(fmt.Sprintf("%v", val))
			if (offset+1)%grid.dimensions[1] == 0 {
				sb.WriteString("\n")
			}
		}
		return sb.String()
	default:
		return fmt.Sprintf("Cannot print %d-dimensional grid", len(grid.dimensions))
	}
	// be mindful of dimensions
	// maybe redo this with reader/writer
}
