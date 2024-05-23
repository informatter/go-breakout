package main


import (
	"github.com/gdamore/tcell/v2"
)

func renderGameObject(screen tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	// gotten from: https://github.com/gdamore/tcell/blob/main/TUTORIAL.md
	row := y1
	col := x1

	for _, r := range []rune(text) {
		screen.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}