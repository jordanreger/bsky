package util

import "strings"

func ReplaceTag(tag string) string {
	return strings.Replace(tag, "#", "%23", -1)
}
