package ansigo

import (
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
)

// TerminalDimensions returns the the columns and rows of the current active window.
func TerminalDimensions() (int, int, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}

	return int(ws.Col), int(ws.Row), nil
}

// EchoStdin prints user input.
func EchoStdin() {
	exec.Command("stty", "-f", "/dev/tty", "echo").Run()
}

// GetChar unbuffers stdin and stops the printing of user stdin input.
func GetChar(stdin chan string) {
	unbufferStdin()
	unechoStdin()
	b := make([]byte, 3)
	for {
		os.Stdin.Read(b)
		stdin <- key(b)
	}
}

// RestoreTerminalSettings undos the effects of GetChar as well as other window
// modifications caused by various functions throughout this library such as CursorHide
func RestoreTerminalSettings() {
	exec.Command("stty", "sane").Run()
}

/*
PRIVATE
*/

func unbufferStdin() {
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
}

func unechoStdin() {
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
}

func key(b []byte) string {
	if b[0] == 27 {
		switch b[len(b)-1] {
		case 0:
			return "<ESC>"
		case 65:
			return "<Up>"
		case 66:
			return "<Down>"
		case 67:
			return "<Right>"
		case 68:
			return "<Left>"
		}
	} else if b[0] == 10 {
		return "<Enter>"
	}

	return string(b[0])
}
