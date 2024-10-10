package clr

const (
	bold      = "\033[1m"
	italic    = "\033[3m"
	underline = "\033[4m"
)

// Bold wraps the text in bold style.
func Bold(text string) string {
	return WrapWithColor(text, bold)
}

// Italic wraps the text in italic style.
func Italic(text string) string {
	return WrapWithColor(text, italic)
}

// Underline wraps the text in underline style.
func Underline(text string) string {
	return WrapWithColor(text, underline)
}
