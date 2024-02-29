package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/jordanreger/htmlsky/api"
	"github.com/jordanreger/htmlsky/types"
)

func GetActorPage(actor types.Actor) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	actor.Feed = api.GetActorFeed(actor)
	var actor_page bytes.Buffer
	err := t.ExecuteTemplate(&actor_page, "actor.html", actor)
	if err != nil {
		fmt.Println(err)
	}
	return actor_page.String()
}
