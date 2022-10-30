package str

import "unicode/utf8"

func Substr(s string, start, length int) string {
	strLen := utf8.RuneCountInString(s)
	if strLen <= 0 {
		return ""
	}
	if length > strLen {
		length = strLen
	}

	runes := []rune(s)

	return string(runes[start:length])
}

func Limit(s string, start, length int, append string) string {
	strLen := utf8.RuneCountInString(s)
	if strLen <= 0 {
		return ""
	}

	if length >= strLen {
		length = strLen
		append = ""
	}

	runes := []rune(s)

	return string(runes[start:length]) + append
}
