package view

import (
	"os"

	"zedterm/doc_status"
	"zedterm/terminal"
	"zedterm/utils"

	"github.com/nsf/termbox-go"
)

type View struct {
	TextLocX     int
	TextLocY     int
	TotalLineNum int

	ScrollOffsetX int
	ScrollOffsetY int

	buffer *buffer

	hint string
}

func NewView() *View {
	v := &View{
		buffer: newBuffer(),
	}
	if len(os.Args) > 1 {
		// TODO: check whether load file success
		_ = v.loadFile(os.Args[1])
	}
	return v
}

func (v *View) SetHint(hint string) {
	v.hint = hint
}

func (v *View) SetDefaultHint() {
	v.SetHint("")
}

func (v *View) SaveFile() error {
	return v.buffer.saveFile()
}

func (v *View) loadFile(fileName string) error {
	return v.buffer.loadFile(fileName)
}

func (v *View) getVisibleText(s string, from, to int) string {
	if from > to {
		return ""
	}
	if from >= len(s) {
		return ""
	}
	return s[from:utils.MinInt(len(s), to)]
}

func (v *View) size() (width, height int) {
	width, height = terminal.Size()
	height -= 1
	return
}

func (v *View) InsertChar(ch rune) {
	v.buffer.insertChar(v.TextLocY, v.TextLocX, ch)
	v.MoveCursor(termbox.KeyArrowRight)
}

func (v *View) Backspace() {
	atRow, atCol := v.TextLocY, v.TextLocX
	if atRow >= 0 || atCol >= 0 {
		v.MoveCursor(termbox.KeyArrowLeft)
		atRow, atCol = v.TextLocY, v.TextLocX
		v.buffer.Delete(atRow, atCol)
	}
}

func (v *View) Delete() {
	atRow, atCol := v.TextLocY, v.TextLocX
	v.buffer.Delete(atRow, atCol)
}

func (v *View) GetDocStatus() *doc_status.DocStatus {
	return &doc_status.DocStatus{
		TextLocX:     v.TextLocX,
		TextLocY:     v.TextLocY,
		TotalLineNum: v.buffer.len(),
		FileName:     v.buffer.fileName,
		Hint:         v.hint,
	}
}

func (v *View) CursorPos() (int, int) {
	return v.TextLocX - v.ScrollOffsetX, v.TextLocY - v.ScrollOffsetY
}

func (v *View) Render() {
	windowWidth, windowHeight := v.size()
	topYIndex := v.ScrollOffsetY
	for i := 0; i < windowHeight; i++ {
		if i+topYIndex < v.buffer.len() {
			from, to := v.ScrollOffsetX, v.ScrollOffsetX+windowWidth
			substr := v.getVisibleText(v.buffer.getLine(i+topYIndex), from, to)
			terminal.PrintLine(i, substr)
		} else {
			terminal.PrintLine(i, "~")
		}
	}
}

func (v *View) NewLine() {
	v.buffer.NewLine(v.TextLocY, v.TextLocX)
	v.MoveCursor(termbox.KeyArrowRight)
}

func (v *View) MoveCursor(key termbox.Key) {
	bufLen := v.buffer.len()
	x := v.TextLocX
	y := v.TextLocY
	switch key {
	case termbox.KeyArrowUp:
		if y > 0 {
			y -= 1
			x = utils.MaxInt(0, utils.MinInt(x, len(v.buffer.getLine(y))-1))
		}
	case termbox.KeyArrowDown:
		if y < bufLen {
			y += 1
			if y == bufLen {
				x = 0
			} else if line := v.buffer.getLine(y); x >= len(line) {
				x = utils.MaxInt(0, len(line)-1)
			}
		}
	case termbox.KeyArrowLeft:
		if x > 0 {
			x -= 1
		} else if y > 0 {
			y -= 1
			x = utils.MaxInt(x, len(v.buffer.getLine(y)))
		}
	case termbox.KeyArrowRight:
		if y < bufLen && x < len(v.buffer.getLine(y)) {
			x += 1
		} else if y <= bufLen-1 {
			y += 1
			x = 0
		}
	}
	v.TextLocX = x
	v.TextLocY = y

	// scroll to move text into view
	windowWidth, windowHeight := v.size()
	// vertically
	if y < v.ScrollOffsetY {
		v.ScrollOffsetY = y
	} else if y >= v.ScrollOffsetY+windowHeight {
		v.ScrollOffsetY = y - windowHeight + 1
	}

	// horizontally
	if x < v.ScrollOffsetX {
		v.ScrollOffsetX = x
	} else if x >= v.ScrollOffsetX+windowWidth {
		v.ScrollOffsetX = x - windowWidth + 1
	}
}
