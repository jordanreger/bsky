package types

import (
	"html/template"
	"time"
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
	Text      string    `json:"text"`
	Type      string    `json:"$type"`
	Langs     []string  `json:"langs"`
	Facets    []Facet   `json:"Facets"`
	Reply     Reply     `json:"reply"`
	CreatedAt time.Time `json:"createdAt"`
	HTML      template.HTML
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
