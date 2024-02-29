package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/jordanreger/htmlsky/types"
)

func GetThreadPage(thread types.Thread) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	var thread_page bytes.Buffer

	err := t.ExecuteTemplate(&thread_page, "thread.html", thread)
	if err != nil {
		fmt.Println(err)
	}
	return thread_page.String()
}
