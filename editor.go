package main

import (
	"zedterm/terminal"
	"zedterm/ui/view"

	"github.com/nsf/termbox-go"
)

type editor struct {
	view *view.View

	needQuit bool
}

func newEditor() *editor {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	return &editor{
		view: view.NewView(),
	}
}

func (e *editor) run() {
	for {
		terminal.HideCursor()
		e.view.Render()
		terminal.ShowCursor(e.view.CursorPos())
		terminal.Flush()
		if e.needQuit {
			terminal.Terminate()
			break
		}
		ev := termbox.PollEvent()
		e.processEvent(ev)
	}
}

func (e *editor) processEvent(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyCtrlQ:
		e.needQuit = true
	case termbox.KeyArrowUp, termbox.KeyArrowDown, termbox.KeyArrowLeft, termbox.KeyArrowRight:
		e.view.MoveCursor(ev.Key)
	}
}
