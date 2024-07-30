package view

import (
	"os"

	"zedterm/doc_status"
	"zedterm/terminal"
	"zedterm/utils"

	"github.com/nsf/termbox-go"
)

type View struct {
	TextLocX int
	TextLocY int

	buffer *buffer
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

func (v *View) GetDocStatus() *doc_status.DocStatus {
	return &doc_status.DocStatus{
		TextLocX: v.TextLocX,
		TextLocY: v.TextLocY,
		FileName: v.buffer.fileName,
	}
}

func (v *View) CursorPos() (int, int) {
	return v.TextLocX, v.TextLocY
}

func (v *View) Render() {
	width, height := v.size()
	for i := 0; i < height; i++ {
		if i < len(v.buffer.lines) {
			from, to := 0, width
			substr := v.getVisibleText(v.buffer.lines[i], from, to)
			terminal.PrintLine(i, substr)
		} else {
			terminal.PrintLine(i, "~")
		}
	}
}

func (v *View) MoveCursor(key termbox.Key) {
	_, windowHeight := v.size()
	x, y := v.TextLocX, v.TextLocY
	switch key {
	case termbox.KeyArrowUp:
		if y > 0 {
			y -= 1
			x = utils.MinInt(x, len(v.buffer.lines[y]))
		}
	case termbox.KeyArrowDown:
		if y < len(v.buffer.lines)-1 {
			y += 1
			x = utils.MinInt(x, len(v.buffer.lines[y]))
		}
	case termbox.KeyArrowLeft:
		if x > 0 {
			x -= 1
		} else if y > 0 {
			y -= 1
			x = len(v.buffer.lines[y])
		}
	case termbox.KeyArrowRight:
		if x < len(v.buffer.lines[y]) {
			x += 1
		} else if y < windowHeight-1 && y < len(v.buffer.lines)-1 {
			y += 1
			x = utils.MinInt(x, len(v.buffer.lines[y]))
		}
	default:
	}
	v.TextLocX, v.TextLocY = x, y
}
