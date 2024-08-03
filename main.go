package main

import (
	"fmt"
	"runtime"

	"zedterm/terminal"

	"github.com/nsf/termbox-go"
)

func main() {
	e := newEditor()
	defer func() {
		if err := recover(); err != nil {
			terminal.Clear()
			terminal.Flush()
			fmt.Println(err)
			buf := make([]byte, 1024000)
			runtime.Stack(buf, false)
			fmt.Println(string(buf))
			termbox.PollEvent()
		}
	}()
	e.run()
}
