package ansigo

import "fmt"

const DISPLAY_ATTR = "\x1B[%vm%v\x1B[0m"

func Bright(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 1, txt)
}

func Underscore(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 4, txt)
}

func Hidden(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 8, txt)
}

func FgBlack(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 30, txt)
}

func FgRed(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 31, txt)
}

func FgGreen(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 32, txt)
}

func FgYellow(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 33, txt)
}

func FgBlue(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 34, txt)
}

func FgMagenta(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 35, txt)
}

func FgCyan(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 36, txt)
}

func FgWhite(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 37, txt)
}

func BgBlack(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 40, txt)
}

func BgRed(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 41, txt)
}

func BgGreen(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 42, txt)
}

func BgYellow(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 43, txt)
}

func BgBlue(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 44, txt)
}

func BgMagenta(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 45, txt)
}

func BgCyan(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 46, txt)
}

func BgWhite(txt string) string {
	return fmt.Sprintf(DISPLAY_ATTR, 47, txt)
}
