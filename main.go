package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/jordanreger/htmlsky/util"
)

//go:embed all:public
var publicFiles embed.FS
var publicFS = fs.FS(publicFiles)
var public, _ = fs.Sub(publicFS, "public")

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/raw/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	mux.HandleFunc("/embed/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// redirect if {handle} is empty
	mux.HandleFunc("/profile/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	mux.HandleFunc("/raw/profile/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	mux.HandleFunc("/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := getActorProfile(did)
		page := getActorPage(actor)

		fmt.Fprint(w, page)
	})
	mux.HandleFunc("/raw/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := getActorProfile(did)
		actor.Feed = getActorFeed(actor)
		res, _ := json.MarshalIndent(actor, "", "    ")

		fmt.Fprint(w, string(res))
	})
	mux.HandleFunc("/embed/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := getActorProfile(did)
		feed := getActorFeed(actor)
		page := getActorPageEmbed(actor, feed)

		fmt.Fprint(w, page)
	})

	// redirect if {post} is empty
	mux.HandleFunc("/profile/{handle}/{post}/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	mux.HandleFunc("/raw/profile/{handle}/post/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	mux.HandleFunc("/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		thread := getThread(at_uri)
		page := getThreadPage(thread)

		fmt.Fprint(w, page)
	})

	mux.HandleFunc("/raw/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		res, _ := json.MarshalIndent(getThread(at_uri), "", "    ")

		fmt.Fprint(w, string(res))
	})

	mux.HandleFunc("/embed/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		thread := getThread(at_uri)
		page := getThreadPageEmbed(thread)

		fmt.Fprint(w, page)
	})

	mux.Handle("/", http.FileServer(http.FS(public)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
