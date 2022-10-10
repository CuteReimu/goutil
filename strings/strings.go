package strings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsEmpty 判断字符串是否为空
func IsEmpty(s string) bool {
	return len(s) == 0
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// IsBlank 判断字符串是否全是空白
func IsBlank(s string) bool {
	// Fast path for ASCII: look for the first ASCII non-space byte
	start := 0
	for ; start < len(s); start++ {
		c := s[start]
		if c >= utf8.RuneSelf {
			// If we run into a non-ASCII byte, fall back to the
			// slower unicode-aware method on the remaining bytes
			return len(strings.TrimLeftFunc(s[start:], unicode.IsSpace)) == 0
		}
		if asciiSpace[c] == 0 {
			break
		}
	}
	return start == len(s)
}
