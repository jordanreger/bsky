package bsky

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
