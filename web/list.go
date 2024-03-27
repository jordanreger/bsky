package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/jordanreger/htmlsky/bsky"
)

func GetListPage(list bsky.List) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	var list_page bytes.Buffer
	err := t.ExecuteTemplate(&list_page, "list.html", list)
	if err != nil {
		fmt.Println(err)
	}
	return list_page.String()
}
