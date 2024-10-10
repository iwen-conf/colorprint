package clr

var Reset = "\033[0m"

// WrapWithColor wraps the text with the given color code.
func WrapWithColor(text, color string) string {
	return color + text + Reset
}
