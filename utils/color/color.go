package color

import (
	"fmt"
)

const colorFmt = "\x1b[%dm%s\x1b[0m"

type Color int

const (
	Black     Color = iota + 30 // 30
	Red                         // 31
	Green                       // 32
	Yellow                      // 33
	Blue                        // 34
	Magenta                     // 35
	Cyan                        // 36
	LightGray                   // 37
	DarkGray  = 90

	Bold = 1
)

func Colorize(s string, c Color) string {
	return fmt.Sprintf(colorFmt, c, s)
}
