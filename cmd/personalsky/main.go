package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/jordanreger/htmlsky/api"
	"github.com/jordanreger/htmlsky/util"
)

/* SET THIS BEFORE BUILDING */
var handle = "did:plc:27rjcwbur2bizjjx3zakeme5"

//go:embed all:public
var publicFiles embed.FS
var publicFS = fs.FS(publicFiles)
var public, _ = fs.Sub(publicFS, "public")

func main() {
	mux := http.NewServeMux()

	/* REDIRECTS */

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.ServeFileFS(w, r, public, r.URL.Path)
		}
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

	log.Fatal(http.ListenAndServe(":8081", mux))
}
