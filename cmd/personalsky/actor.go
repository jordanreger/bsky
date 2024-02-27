package main

import (
	"bytes"
	"html/template"

	"github.com/jordanreger/htmlsky/api"
	"github.com/jordanreger/htmlsky/types"
)

func GetActorPage(actor types.Actor) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	actor.Feed = api.GetActorFeed(actor)
	var actor_page bytes.Buffer
	t.ExecuteTemplate(&actor_page, "actor.html", actor)
	return actor_page.String()
}
