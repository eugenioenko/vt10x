package vt10x

// ANSI color values
const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGrey
	DarkGrey
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

const (
	colorTrueColor Color = 1 << 25
)

// Default colors are potentially distinct to allow for special behavior.
// For example, a transparent background. Otherwise, the simple case is to
// map default colors to another color.
const (
	DefaultFG Color = 1<<24 + iota
	DefaultBG
	DefaultCursor
)

// Color maps to the ANSI colors [0, 16) and the xterm colors [16, 256).
// True color (24-bit RGB) values have the colorTrueColor flag set.
type Color uint32

// ANSI returns true if Color is within [0, 16).
func (c Color) ANSI() bool {
	return c < 16 && c&colorTrueColor == 0
}

// TrueColor returns true if Color represents a 24-bit RGB value.
func (c Color) TrueColor() bool {
	return c&colorTrueColor != 0
}

// RGB extracts the red, green, and blue components from a true color value.
func (c Color) RGB() (r, g, b uint8) {
	v := uint32(c &^ colorTrueColor)
	return uint8(v >> 16), uint8(v >> 8), uint8(v)
}

// TrueColorValue creates a Color from RGB components with the true color flag.
func TrueColorValue(r, g, b uint8) Color {
	return Color(int(r)<<16|int(g)<<8|int(b)) | colorTrueColor
}
