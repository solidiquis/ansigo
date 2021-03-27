package ansigo

import (
	"golang.org/x/sys/unix"
	"os"
	"os/exec"
)

func TerminalDimensions() (int, int, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}

	return int(ws.Col), int(ws.Row), nil
}

func UnbufferStdin() {
	exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
}

func UnechoStdin() {
	exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
}

func EchoStdin() {
	exec.Command("stty", "-f", "/dev/tty", "echo").Run()
}

func GetChar(stdin chan string) {
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		stdin <- string(b)
	}
}

func RestoreTerminalSettings() {
	exec.Command("stty", "sane").Run()
}
