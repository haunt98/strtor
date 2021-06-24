package normalize

import (
	"regexp"
	"strings"
)

const (
	normalizedChar = '_'
)

var reNormalizedDuplicate = regexp.MustCompile(`_+`)

func Normalize(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	// whitelist approach
	tempText := ""
	for _, c := range text {
		if isValidChar(c) {
			tempText += string(c)
			continue
		}

		tempText += string(normalizedChar)
	}
	text = tempText

	// remove duplicate _
	// a__b -> a_b
	text = reNormalizedDuplicate.ReplaceAllString(text, string(normalizedChar))

	// remove prefix, postfix _
	// _abc_ -> abc
	text = strings.TrimPrefix(text, string(normalizedChar))
	text = strings.TrimSuffix(text, string(normalizedChar))

	return text
}

func isValidText(text string) bool {
	for _, c := range text {
		if !isValidChar(c) {
			return false
		}
	}

	return true
}

func isValidChar(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9') ||
		c == normalizedChar
}
