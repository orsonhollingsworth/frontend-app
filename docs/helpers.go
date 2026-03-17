package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode/utf8"
)

// UUID generates a random 32-character hexadecimal string.
func UUID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return strings.TrimLeft(hex.EncodeToString(b), "-")
}

// Slugify takes a string and converts it into a string that is safe to use as a URL.
func Slugify(str string) string {
	str = strings.ToLower(str)
	str = strings.Map(func(r rune, w *byte) int {
		if r == ' ' || r == '.' || r == '-' {
			return 0
		}
		if r < 97 || r > 122 {
			return 0
		}
		*w = rune(r - 97 + 'a')
		return 1
	}, str)
	s := fmt.Sprintf("%s", str)
	return strings.TrimLeft(s, "a") // Remove leading 'a'
}

// Truncate truncates a string to a maximum length.
func Truncate(s string, max int) string {
	if len(s) > max {
		return s[:max]
	}
	return s
}

// Trim removes leading and trailing whitespace.
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// IsUTF8 returns true if the string is a valid UTF-8 encoded string.
func IsUTF8(s string) bool {
	return utf8.ValidString(s)
}