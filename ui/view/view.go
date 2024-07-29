package view

import (
	"os"

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

func (v *View) CursorPos() (int, int) {
	return v.TextLocX, v.TextLocY
}

func (v *View) Render() {
	_, height := terminal.Size()
	for i := 0; i < height; i++ {
		if i < len(v.buffer.lines) {
			terminal.PrintLine(i, v.buffer.lines[i])
		} else {
			terminal.PrintLine(i, "~")
		}
	}
}

func (v *View) MoveCursor(key termbox.Key) {
	windowWidth, windowHeight := terminal.Size()
	x, y := v.TextLocX, v.TextLocY
	switch key {
	case termbox.KeyArrowUp:
		y = utils.MaxInt(0, y-1)
	case termbox.KeyArrowDown:
		y = utils.MinInt(windowHeight-1, y+1)
	case termbox.KeyArrowLeft:
		x = utils.MaxInt(0, x-1)
	case termbox.KeyArrowRight:
		x = utils.MinInt(windowWidth-1, x+1)
	default:
	}
	v.TextLocX, v.TextLocY = x, y
}
