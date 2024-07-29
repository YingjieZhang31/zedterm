package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	for {
		ev := termbox.PollEvent()
		if ch := ev.Ch; ch != 0 {
			fmt.Printf("Entered: %c\n", ch)
			if ch == 'q' {
				break
			}
		}
	}
}
