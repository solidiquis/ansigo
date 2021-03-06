# ansigo
A library to make the controlling of ANSI terminal emulators' fonts, colors, and cursors, a little less painful — also can be used to detect key presses in the terminal.

Documentation can be found at https://pkg.go.dev/github.com/solidiquis/ansigo

## Installation
```
$ go get github.com/solidiquis/ansigo
```

## Demo
Moving the cursor around the screen Vim-style:
```go
package main

import (
	ansi "github.com/solidiquis/ansigo"
)

func main() {
	ansi.EraseScreen()
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
```
<img src="https://github.com/solidiquis/solidiquis/blob/master/assets/ansigo_demo.gif?raw=true">
