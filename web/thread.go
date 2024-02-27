package main

import (
	"bytes"
	"html/template"

	"github.com/jordanreger/htmlsky/types"
)

func GetThreadPage(thread types.Thread) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	var thread_page bytes.Buffer
	t.ExecuteTemplate(&thread_page, "thread.html", thread)
	return thread_page.String()
}
