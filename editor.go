package main

import (
	"zedterm/terminal"

	"github.com/nsf/termbox-go"
)

type Editor struct {
	needQuit bool
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
		terminal.ShowCursor(0, 0)
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
	if ev.Key == termbox.KeyCtrlQ {
		e.needQuit = true
		return
	}
}

func (e *Editor) refreshScreen() {
	_, height := terminal.Size()
	for i := 0; i < height; i++ {
		terminal.PrintLine(i, "~")
	}
}
