package ansigo

import "fmt"

func CursorUp(n int) {
	fmt.Printf("\x1B[%dA", n)
}

func CursorDown(n int) {
	fmt.Printf("\x1B[%dB", n)
}

func CursorForward(n int) {
	fmt.Printf("\x1B[%dC", n)
}

func CursorBackward(n int) {
	fmt.Printf("\x1B[%dD", n)
}

func CursorHome() {
	fmt.Print("\x1B[H")
}

func CursorSetPos(row, col int) {
	fmt.Printf("\x1B[%d;%dH", row, col)
}

func CursorHide() {
	fmt.Print("\x1B[?25l")
}

func CursorShow() {
	fmt.Print("\x1B[?25h")
}

func CursorSavePos() {
	fmt.Print("\x1B[s")
}

func CursorRestorePos() {
	fmt.Print("\x1B[u")
}
