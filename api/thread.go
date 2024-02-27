package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/jordanreger/htmlsky/types"
	"github.com/jordanreger/htmlsky/util"
)

type t_res struct {
	Thread types.Thread
}

func GetThread(at_uri string) types.Thread {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.feed.getPostThread?uri=" + at_uri)
	if err != nil {
		fmt.Println(err)
	}

	var t_body t_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &t_body)

	thread := t_body.Thread

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		thread.Post.RKey = util.GetRKey(thread.Post.URI)
	}()
	/*
		wg.Add(1)
		go func() {
			defer wg.Done()
			thread.Post.Author = getActorProfile(thread.Post.Author.DID)
		}()
	*/
	wg.Add(1)
	go func() {
		defer wg.Done()
		thread.Post.Record.HTML = util.FacetsToHTML(thread.Post.Record.Text, thread.Post.Record.Facets)
	}()

	wg.Wait()

	for i := range thread.Replies {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			thread.Replies[i].Post.RKey = util.GetRKey(thread.Replies[i].Post.URI)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if thread.Replies[i].Post.Author.DisplayName == "" {
				thread.Replies[i].Post.Author.DisplayName = thread.Replies[i].Post.Author.Handle
			}
			if thread.Replies[i].Post.Author.Avatar == "" {
				thread.Replies[i].Post.Author.Avatar = "/avatar.jpeg"
			}
			if thread.Replies[i].Post.Author.Banner == "" {
				thread.Replies[i].Post.Author.Banner = "/banner.jpeg"
			}
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			thread.Replies[i].Post.Record.HTML = util.FacetsToHTML(thread.Replies[i].Post.Record.Text, thread.Replies[i].Post.Record.Facets)
		}(i)

	}
	wg.Wait()

	return thread
}
