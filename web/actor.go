package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/jordanreger/htmlsky/bsky"
)

func GetActorPage(actor bsky.Actor) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	var actor_page bytes.Buffer
	err := t.ExecuteTemplate(&actor_page, "actor.html", actor)
	if err != nil {
		fmt.Println(err)
	}
	return actor_page.String()
}
