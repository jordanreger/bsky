package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jordanreger/htmlsky/types"
	"github.com/jordanreger/htmlsky/util"
)

func GetActorProfile(did string) types.Actor {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.actor.getProfile?actor=" + did)
	if err != nil {
		fmt.Println(err)
	}

	var actor types.Actor
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

	actor.Feed = GetActorFeed(actor)

	descriptionFacets := util.ParseFacets(actor.Description)
	actor.DescriptionHTML = util.FacetsToHTML(actor.Description, descriptionFacets)

	return actor
}
