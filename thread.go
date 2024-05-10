package bsky

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type t_res struct {
	Thread Thread
}

func GetThread(at_uri string) Thread {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.feed.getPostThread?uri=" + at_uri)
	if err != nil {
		fmt.Println(err)
	}

	var t_body t_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &t_body)

	thread := t_body.Thread
	thread.Post.RKey()

	return thread
}
