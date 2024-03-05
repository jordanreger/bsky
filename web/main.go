package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/jordanreger/htmlsky/bsky"
	"github.com/jordanreger/htmlsky/util"
)

//go:embed all:public
var publicFiles embed.FS
var publicFS = fs.FS(publicFiles)
var public, _ = fs.Sub(publicFS, "public")

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.ServeFileFS(w, r, public, r.URL.Path)
			return
		}
		handle := "htmlsky.app"
		did := util.GetDID(handle)
		actor := bsky.GetActorProfile(did)
		page := GetActorPage(actor)

		fmt.Fprint(w, page)
	})

	/* REDIRECTS */
	mux.HandleFunc("/raw/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Usage: /raw/profile/{handle}[/post/{rkey}]")
	})
	mux.HandleFunc("/raw/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/raw/", http.StatusSeeOther)
	})
	mux.HandleFunc("/raw/profile/{handle}/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/raw/", http.StatusSeeOther)
	})
	mux.HandleFunc("/embed/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Usage: /embed/profile/{handle}[/post/{rkey}]")
	})
	mux.HandleFunc("/embed/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/embed/", http.StatusSeeOther)
	})
	mux.HandleFunc("/profile/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	mux.HandleFunc("/profile/{handle}/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	/* ROUTES */

	// actor
	mux.HandleFunc("/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := bsky.GetActorProfile(did)
		page := GetActorPage(actor)

		fmt.Fprint(w, page)
	})
	mux.HandleFunc("/raw/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := bsky.GetActorProfile(did)
		res, _ := json.MarshalIndent(actor, "", "    ")

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(res))
	})
	mux.HandleFunc("/embed/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := bsky.GetActorProfile(did)
		page := GetActorPageEmbed(actor)

		fmt.Fprint(w, page)
	})

	// thread
	mux.HandleFunc("/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		thread := bsky.GetThread(at_uri)
		page := GetThreadPage(thread)

		fmt.Fprint(w, page)
	})

	mux.HandleFunc("/raw/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		res, _ := json.MarshalIndent(bsky.GetThread(at_uri), "", "    ")

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, string(res))
	})

	mux.HandleFunc("/embed/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		thread := bsky.GetThread(at_uri)
		page := GetThreadPageEmbed(thread)

		fmt.Fprint(w, page)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
