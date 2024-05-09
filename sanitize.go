package util

import (
	"html/template"
	"strings"
)

func Sanitize(text string) string {
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	return text
}

func ReplaceNewlines(text string) template.HTML {
	text = strings.Replace(text, "\n", "<br>", -1)
	return template.HTML(text)
}
