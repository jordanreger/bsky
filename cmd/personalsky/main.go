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

var handle = "did:plc:27rjcwbur2bizjjx3zakeme5"

//go:embed all:public
var publicFiles embed.FS
var publicFS = fs.FS(publicFiles)
var public, _ = fs.Sub(publicFS, "public")

func main() {
	mux := http.NewServeMux()

	/* REDIRECTS */
	// redirect if {post} is empty
	mux.HandleFunc("/{post}/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	/* ROUTES */

	mux.HandleFunc("/", func(w http.ResponseWriter r *http.Request) {
		did := util.GetDID(handle)
		actor := api.GetActorProfile(did)
		page := GetActorPage(actor)

		fmt.Fprint(w, page)
	})

	// thread
	mux.HandleFunc("/post/{rkey}/", func(w http.ResponseWriter, r *http.Request) {
		rkey := r.PathValue("rkey")

		did := util.GetDID(handle)
		at_uri := util.GetPostURI(did, rkey)
		thread := api.GetThread(at_uri)
		page := GetThreadPage(thread)

		fmt.Fprint(w, page)
	})

	// static
	mux.Handle("/", http.FileServer(http.FS(public)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
