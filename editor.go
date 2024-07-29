package main

import (
	"zedterm/terminal"

	"github.com/nsf/termbox-go"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Editor struct {
	needQuit bool

	TextLocX int
	TextLocY int
}

func newEditor() *Editor {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	return &Editor{}
}

func (e *Editor) run() {
	for {
		terminal.HideCursor()
		e.refreshScreen()
		terminal.ShowCursor(e.TextLocX, e.TextLocY)
		terminal.Flush()
		if e.needQuit {
			terminal.Terminate()
			break
		}
		ev := termbox.PollEvent()
		e.processEvent(ev)
	}
}

func (e *Editor) processEvent(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyCtrlQ:
		e.needQuit = true
	case termbox.KeyArrowUp, termbox.KeyArrowDown, termbox.KeyArrowLeft, termbox.KeyArrowRight:
		e.moveCursor(ev.Key)
	}
}

func (e *Editor) moveCursor(key termbox.Key) {
	windowWidth, windowHeight := terminal.Size()
	x, y := e.TextLocX, e.TextLocY
	switch key {
	case termbox.KeyArrowUp:
		y = max(0, y-1)
	case termbox.KeyArrowDown:
		y = min(windowHeight-1, y+1)
	case termbox.KeyArrowLeft:
		x = max(0, x-1)
	case termbox.KeyArrowRight:
		x = min(windowWidth-1, x+1)
	default:
	}
	e.TextLocX, e.TextLocY = x, y
}

func (e *Editor) refreshScreen() {
	_, height := terminal.Size()
	for i := 0; i < height; i++ {
		terminal.PrintLine(i, "~")
	}
}
