package grid2

import (
	"strings"
)

//type GridPrintOptions struct {
//	// TODO: add support for labelling rows and columns with ASCII art
//	LabelColumns bool
//	LabelRows    bool
//	// Compact mode; one cell == one character
//	Compact bool
//}

func StringifyRuneGrid(g *Grid[rune]) string {
	var sb strings.Builder
	for row := range g.Rows() {
		for _, r := range row {
			sb.WriteRune(r)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
