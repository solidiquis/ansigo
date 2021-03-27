package ansigo

import "fmt"

const (
	SCREEN_ESC = "\x1B[%vJ"
	LINE_ESC   = "\x1B[%vK"
)

func EraseToEndln() {
	fmt.Printf(LINE_ESC, "")
}

func EraseToStartln() {
	fmt.Printf(LINE_ESC, 1)
}

func EraseLine() {
	fmt.Printf(LINE_ESC, 2)
}

func EraseDown() {
	fmt.Printf(SCREEN_ESC, "")
}

func EraseUp() {
	fmt.Printf(SCREEN_ESC, 1)
}

func EraseScreen() {
	fmt.Printf(SCREEN_ESC, 2)
}

func Backspace() {
	fmt.Print("\b \b")
}
