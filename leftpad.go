package leftpad

import (
	"strings"
	"unicode/utf8"

	"github.com/seuz113/leftpad/driver/cycle"
)

var default_char = ''

// Comments space.
// testing comments
func Format(s string, size int) string {
	return FormatRune(s, size, cycle.DEFAULT_CHAR)
}

// Comments space.
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
