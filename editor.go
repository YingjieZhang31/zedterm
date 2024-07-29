package main

import (
	"fmt"

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
		if e.needQuit {
			termbox.Close()
			break
		}
		e.refreshScreen()
		ev := termbox.PollEvent()
		e.processEvent(ev)
	}
}

func (e *Editor) processEvent(ev termbox.Event) {
	if ch := ev.Ch; ch != 0 {
		if ch == 'q' {
			e.needQuit = true
		}
	}
}

func (e *Editor) refreshScreen() {
	_, height := termbox.Size()
	for i := 0; i < height; i++ {
		fmt.Printf("~\r\n")
	}
}
