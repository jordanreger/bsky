package bsky

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

type Actor struct {
	DID            string        `json:"did"`
	Handle         string        `json:"handle"`
	DisplayName    string        `json:"displayName"`
	Description    string        `json:"description"`
	Avatar         string        `json:"avatar"`
	Banner         string        `json:"banner"`
	FollowsCount   int           `json:"followsCount"`
	FollowersCount int           `json:"followersCount"`
	PostsCount     int           `json:"postsCount"`
	IndexedAt      time.Time     `json:"indexedAt"`
	Labels         []interface{} `json:"labels"`
}

func (actor Actor) Feed() Feed {
	return GetActorFeed(actor)
}

func (actor Actor) DescriptionHTML() template.HTML {
	return template.HTML(actor.Description)
}

func GetActorProfile(did string) Actor {
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

	return actor
}
