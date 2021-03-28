package main

import (
	ansi "github.com/solidiquis/ansigo"
)

func main() {
	ansi.EraseScreen()
	ansi.UnbufferStdin()
	ansi.UnechoStdin()
	ansi.CursorSetPos(0, 0)

	stdin := make(chan string, 1)
	go ansi.GetChar(stdin)

	for {
		switch <-stdin {
		case "h":
			ansi.CursorBackward(1)
		case "j":
			ansi.CursorDown(1)
		case "k":
			ansi.CursorUp(1)
		case "l":
			ansi.CursorForward(1)
		}
	}
}
