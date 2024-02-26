package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/jordanreger/htmlsky/util"
)

type Actor struct {
	DID             string `json:"did"`
	Handle          string `json:"handle"`
	DisplayName     string `json:"displayName"`
	Description     string `json:"description"`
	Avatar          string `json:"avatar"`
	Banner          string `json:"banner"`
	FollowersCount  int    `json:"followersCount"`
	FollowsCount    int    `json:"followsCount"`
	PostsCount      int    `json:"postsCount"`
	Feed            []FeedItem
	DescriptionHTML template.HTML
}

func getActorProfile(did string) Actor {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.actor.getProfile?actor=" + did)
	if err != nil {
		fmt.Println(err)
	}

	var actor Actor
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &actor)

	if actor.DisplayName == "" {
		actor.DisplayName = actor.Handle
	}
	if actor.Avatar == "" {
		actor.Avatar = "/avatar.jpeg"
	}
	if actor.Banner == "" {
		actor.Banner = "/banner.jpeg"
	}

	actor.Feed = getActorFeed(actor)

	descriptionFacets := util.ParseFacets(actor.Description)
	actor.DescriptionHTML = util.FacetsToHTML(actor.Description, descriptionFacets)

	return actor
}

func getActorPage(actor Actor) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	actor.Feed = getActorFeed(actor)
	var actor_page bytes.Buffer
	t.ExecuteTemplate(&actor_page, "actor.html", actor)
	return actor_page.String()
}
