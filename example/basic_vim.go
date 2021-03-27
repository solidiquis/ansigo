package main

import (
	"fmt"
	ansi "github.com/solidiquis/ansigo"
)

func switchMode(mode *string, newMode string) {
	_, row, err := ansi.TerminalDimensions()
	if err != nil {
		panic(err)
	}
	ansi.CursorSavePos()

	*mode = newMode

	ansi.CursorSetPos(row, 0)
	ansi.EraseLine()
	fmt.Print(ansi.Bright(newMode))
	ansi.CursorRestorePos()
}

func main() {
	_, row, err := ansi.TerminalDimensions()
	if err != nil {
		panic(err)
	}

	mode := "NORMAL"

	ansi.EraseScreen()
	ansi.UnbufferStdin()
	ansi.UnechoStdin()
	ansi.CursorSetPos(row, 0)
	fmt.Print(ansi.Bright("NORMAL"))
	ansi.CursorSetPos(0, 0)

	stdin := make(chan string, 1)
	go ansi.GetChar(stdin)

	for {
		select {
		case ch := <-stdin:
			switch ch[0] {
			case 27:
				switchMode(&mode, "NORMAL")
				continue
			case 105:
				switchMode(&mode, "INSERT")
				continue
			}

			if mode == "NORMAL" {
				switch ch[0] {
				case 104:
					ansi.CursorBackward(1)
				case 106:
					ansi.CursorDown(1)
				case 107:
					ansi.CursorUp(1)
				case 108:
					ansi.CursorForward(1)
				}
				continue
			}

			if mode == "INSERT" {
				fmt.Print(string(ch))
			}
		}
	}
}
