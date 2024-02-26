package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jordanreger/htmlsky/util"
)

type Feed = []FeedItem

type FeedItem struct {
	Post Post `json:"post"`
}

type f_res struct {
	Feed   Feed   `json:"feed"`
	Cursor string `json:"cursor"`
}

func getActorFeed(actor Actor) Feed {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.feed.getAuthorFeed?actor=" + actor.DID)
	if err != nil {
		fmt.Println(err)
	}

	var f_body f_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &f_body)

	feed := f_body.Feed
	for i := range feed {
		feed[i].Post.RKey = util.GetRKey(feed[i].Post.URI)
	}

	return feed
}
