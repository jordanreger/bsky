package main

import (
	"bytes"
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
	Author      Actor     `json:"author"`
	Record      Record    `json:"record"`
	Embed       Embed     `json:"embed"`
	ReplyCount  int       `json:"replyCount"`
	RepostCount int       `json:"repostCount"`
	LikeCount   int       `json:"likeCount"`
	IndexedAt   time.Time `json:"indexedAt"`
	Labels      []string  `json:"labels"`
}

type Record struct {
	Text      string       `json:"text"`
	Type      string       `json:"$type"`
	Langs     []string     `json:"langs"`
	Facets    []util.Facet `json:"Facets"`
	Reply     Reply        `json:"reply"`
	CreatedAt time.Time    `json:"createdAt"`
	HTML      template.HTML
}

type Embed struct {
	Type   string  `json:"$type"`
	Images []Image `json:"images"`
}

type Image struct {
	Thumb    string `json:"thumb"`
	FullSize string `json:"fullsize"`
	Alt      string `json:"alt"`
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

func getThread(at_uri string) Thread {
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
		go func() {
			defer wg.Done()
			thread.Replies[i].Post.RKey = util.GetRKey(thread.Replies[i].Post.URI)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			if thread.Replies[i].Post.Author.DisplayName == "" {
				thread.Replies[i].Post.Author.DisplayName = thread.Replies[i].Post.Author.Handle
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			thread.Replies[i].Post.Record.HTML = util.FacetsToHTML(thread.Replies[i].Post.Record.Text, thread.Replies[i].Post.Record.Facets)
		}()

		wg.Wait()
	}

	return thread
}

func getThreadPage(thread Thread) string {
	t := template.Must(template.ParseFS(publicFiles, "public/*"))
	var thread_page bytes.Buffer
	t.ExecuteTemplate(&thread_page, "thread.html", thread)
	return thread_page.String()
}
