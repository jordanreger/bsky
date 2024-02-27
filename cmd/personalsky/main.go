package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/jordanreger/htmlsky/api"
	"github.com/jordanreger/htmlsky/util"
)

//go:embed all:public
var publicFiles embed.FS
var publicFS = fs.FS(publicFiles)
var public, _ = fs.Sub(publicFS, "public")

func main() {
	mux := http.NewServeMux()

	/* REDIRECTS */

	// redirect if path is just /*/
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
	// redirect if {post} is empty
	mux.HandleFunc("/profile/{handle}/{post}/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	mux.HandleFunc("/raw/profile/{handle}/post/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	/* ROUTES */

	// actor
	mux.HandleFunc("/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := api.GetActorProfile(did)
		page := GetActorPage(actor)

		fmt.Fprint(w, page)
	})
	mux.HandleFunc("/raw/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := api.GetActorProfile(did)
		res, _ := json.MarshalIndent(actor, "", "    ")

		fmt.Fprint(w, string(res))
	})
	mux.HandleFunc("/embed/profile/{handle}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")

		did := util.GetDID(handle)
		actor := api.GetActorProfile(did)
		page := GetActorPageEmbed(actor)

		fmt.Fprint(w, page)
	})

	// thread
	mux.HandleFunc("/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		thread := api.GetThread(at_uri)
		page := GetThreadPage(thread)

		fmt.Fprint(w, page)
	})

	mux.HandleFunc("/raw/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		res, _ := json.MarshalIndent(api.GetThread(at_uri), "", "    ")

		fmt.Fprint(w, string(res))
	})

	mux.HandleFunc("/embed/profile/{handle}/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		handle := r.PathValue("handle")
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		thread := api.GetThread(at_uri)
		page := GetThreadPageEmbed(thread)

		fmt.Fprint(w, page)
	})

	// static
	mux.Handle("/", http.FileServer(http.FS(public)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
