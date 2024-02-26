package main

import (
	"bytes"
	"html/template"
)

func getActorPageEmbed(actor Actor) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	actor.Feed = getActorFeed(actor)
	var actor_page bytes.Buffer
	t.ExecuteTemplate(&actor_page, "actor.embed.html", actor)
	return actor_page.String()
}

func getThreadPageEmbed(thread Thread) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	var thread_page bytes.Buffer
	t.ExecuteTemplate(&thread_page, "thread.embed.html", thread)
	return thread_page.String()
}
