package terminal

import "github.com/nsf/termbox-go"

func Init() {
	termbox.Init()
}

func Terminate() {
	termbox.Close()
}

func Size() (int, int) {
	return termbox.Size()
}

func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func ShowCursor(col, row int) {
	termbox.SetCursor(col, row)
}

func HideCursor() {
	termbox.HideCursor()
}

func PrintLine(row int, text string) {
	fg, bg := termbox.ColorDefault, termbox.ColorDefault
	for i, c := range text {
		termbox.SetCell(i, row, c, fg, bg)
	}
}

func Flush() {
	termbox.Flush()
}
