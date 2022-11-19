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
		for x := grid.offsets[1]; x < grid.offsets[1]+grid.dimensions[1]; x++ {
			if x%10 == 0 {
				sb.WriteString(fmt.Sprintf("%d", IntAbs(x)/10))
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
		for x := grid.offsets[1]; x < grid.offsets[1]+grid.dimensions[1]; x++ {
			sb.WriteString(fmt.Sprintf("%d", IntAbs(x)%10))
		}
		sb.WriteString("\n")
		sb.WriteString(strings.Repeat("-", grid.dimensions[1]))
		sb.WriteString("\n")
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
