package main

import (
	"bytes"
	"html/template"

	"github.com/jordanreger/htmlsky/types"
)

func GetActorPageEmbed(actor types.Actor) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	var actor_page bytes.Buffer
	t.ExecuteTemplate(&actor_page, "actor.embed.html", actor)
	return actor_page.String()
}

func GetThreadPageEmbed(thread types.Thread) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	var thread_page bytes.Buffer
	t.ExecuteTemplate(&thread_page, "thread.embed.html", thread)
	return thread_page.String()
}
