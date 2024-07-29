package view

import (
	"zedterm/terminal"
	"zedterm/utils"

	"github.com/nsf/termbox-go"
)

type View struct {
	TextLocX int
	TextLocY int
}

func NewView() *View {
	return &View{}
}

func (v *View) CursorPos() (int, int) {
	return v.TextLocX, v.TextLocY
}

func (v *View) Render() {
	_, height := terminal.Size()
	terminal.PrintLine(0, "Hello, world!")
	for i := 1; i < height; i++ {
		terminal.PrintLine(i, "~")
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
