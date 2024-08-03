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

	ScrollOffsetX int
	ScrollOffsetY int

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
	return v.TextLocX - v.ScrollOffsetX, v.TextLocY - v.ScrollOffsetY
}

func (v *View) Render() {
	lines := v.buffer.lines
	viewWidth, viewHeight := v.size()
	topYIndex := v.ScrollOffsetY
	for i := 0; i < viewHeight; i++ {
		if i+topYIndex < len(lines) {
			from, to := v.ScrollOffsetX, v.ScrollOffsetX+viewWidth
			substr := v.getVisibleText(lines[i+topYIndex], from, to)
			terminal.PrintLine(i, substr)
		} else {
			terminal.PrintLine(i, "~")
		}
	}
}

func (v *View) MoveCursor(key termbox.Key) {
	lines := v.buffer.lines
	x := v.TextLocX
	y := v.TextLocY
	switch key {
	case termbox.KeyArrowUp:
		if y > 0 {
			y -= 1
			x = utils.MaxInt(0, utils.MinInt(x, len(lines[y])-1))
		}
	case termbox.KeyArrowDown:
		if y < len(lines) {
			y += 1
			if y == len(lines) {
				x = 0
			} else if x >= len(lines[y]) {
				x = utils.MaxInt(0, len(lines[y])-1)
			}
		}
	case termbox.KeyArrowLeft:
		if x > 0 {
			x -= 1
		} else if y > 0 {
			y -= 1
			x = utils.MaxInt(x, len(lines[y]))
		}
	case termbox.KeyArrowRight:
		if y < len(lines) && x < len(lines[y]) {
			x += 1
		} else if y <= len(lines)-1 {
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
