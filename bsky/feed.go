package bsky

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type a_res struct {
	Feed   Feed   `json:"feed"`
	Cursor string `json:"cursor"`
}

func GetActorFeed(actor Actor) Feed {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.feed.getAuthorFeed?actor=" + actor.DID)
	if err != nil {
		fmt.Println(err)
	}

	var f_body a_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &f_body)

	feed := f_body.Feed

	return feed
}

type l_res struct {
	Feed   Feed   `json:"feed"`
	Cursor string `json:"cursor"`
}

/* Must use util.GetListURI() for this */
func GetListFeed(at_uri string) Feed {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.feed.getListFeed?at_uri=" + at_uri)
	if err != nil {
		fmt.Println(err)
	}

	var f_body l_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &f_body)

	feed := f_body.Feed

	return feed
}
