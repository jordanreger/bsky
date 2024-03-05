package bsky

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/jordanreger/htmlsky/util"
)

type Thread struct {
	Type    string   `json:"$type"`
	Post    Post     `json:"post"`
	Replies []Thread `json:"replies"`
}

type Post struct {
	RKey        string
	URI         string    `json:"uri"`
	CID         string    `json:"cid"`
	Author      *Actor    `json:"author"`
	Record      Record    `json:"record"`
	Embed       Embed     `json:"embed"`
	ReplyCount  int       `json:"replyCount"`
	RepostCount int       `json:"repostCount"`
	LikeCount   int       `json:"likeCount"`
	IndexedAt   time.Time `json:"indexedAt"`
	Labels      []string  `json:"labels"`
}

type Record struct {
	Text      string        `json:"text"`
	Type      string        `json:"$type"`
	Langs     []string      `json:"langs"`
	Embed     *Embed        `json:"embed"`
	Facets    []Facet       `json:"facets"`
	Reply     *Reply        `json:"reply"`
	CreatedAt time.Time     `json:"createdAt"`
	HTML      template.HTML `json:"html"`
}

type Reply struct {
	Root   ReplyRoot   `json:"root"`
	Parent ReplyParent `json:"parent"`
}
type ReplyRoot struct {
	CID string `json:"cid"`
	URI string `json:"uri"`
}
type ReplyParent struct {
	CID string `json:"cid"`
	URI string `json:"uri"`
}

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
		thread.Post.Record.HTML = FacetsToHTML(thread.Post.Record.Text, thread.Post.Record.Facets)
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
			thread.Replies[i].Post.Record.HTML = FacetsToHTML(thread.Replies[i].Post.Record.Text, thread.Replies[i].Post.Record.Facets)
		}(i)

	}
	wg.Wait()

	return thread
}
