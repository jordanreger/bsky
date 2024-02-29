package util

import "strings"

func Sanitize(text string) string {
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	return text
}
