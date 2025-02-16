package main

import (
	"zedterm/terminal"
	"zedterm/ui"
	"zedterm/ui/view"

	"github.com/nsf/termbox-go"
)

type editor struct {
	view      *view.View
	statusBar *ui.StatusBar

	needQuit bool
}

func newEditor() *editor {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	return &editor{
		view:      view.NewView(),
		statusBar: ui.NewStatusBar(),
	}
}

func (e *editor) run() {
	for {
		terminal.HideCursor()
		terminal.Clear()
		e.view.Render()
		status := e.view.GetDocStatus()
		e.statusBar.UpdateDocStatus(status)
		e.statusBar.Render()
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
		e.view.SetDefaultHint()
	case termbox.KeyEnter:
		e.view.NewLine()
		e.view.SetDefaultHint()
	case termbox.KeyDelete:
		e.view.Delete()
		e.view.SetDefaultHint()
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		e.view.Backspace()
		e.view.SetDefaultHint()
	case termbox.KeySpace:
		e.view.InsertChar(' ')
	case termbox.KeyCtrlS:
		err := e.view.SaveFile()
		if err != nil {
			e.view.SetHint("Unable to saved!")
		} else {
			e.view.SetHint("Saved!")
		}
	}
	if ev.Ch != 0 {
		e.view.InsertChar(ev.Ch)
	}
}
