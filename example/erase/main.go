package main

import (
	"fmt"
	ansi "github.com/solidiquis/ansigo"
	"time"
)

func main() {
	ansi.EraseScreen()
	ansi.CursorSetPos(0, 0)
	ansi.CursorSavePos()
	ansi.CursorSetPos(5, 5)
	fmt.Print("Hello World")
	ansi.CursorRestorePos()
	ansi.EraseLineSection(5, 11, 16)
	time.Sleep(2 * time.Second)
}
