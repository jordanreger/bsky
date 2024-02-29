package types

import (
	"html/template"
	"time"
)

type Thread struct {
	Type    string   `json:"$type"`
	Post    Post     `json:"post"`
	Replies []Thread `json:"replies,omitempty"`
}

type Post struct {
	RKey        string
	URI         string    `json:"uri"`
	CID         string    `json:"cid"`
	Author      *Actor    `json:"author"`
	Record      Record    `json:"record"`
	Embed       Embed     `json:"embed,omitempty"`
	ReplyCount  int       `json:"replyCount"`
	RepostCount int       `json:"repostCount"`
	LikeCount   int       `json:"likeCount"`
	IndexedAt   time.Time `json:"indexedAt"`
	Labels      []string  `json:"labels,omitempty"`
}

type Record struct {
	Text      string        `json:"text"`
	Type      string        `json:"$type"`
	Langs     []string      `json:"langs,omitempty"`
	Embed     *Embed        `json:"embed,omitempty"`
	Facets    []Facet       `json:"facets,omitempty"`
	Reply     *Reply        `json:"reply,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
	HTML      template.HTML `json:"html,omitempty"`
}

type Reply struct {
	Root   ReplyRoot   `json:"root,omitempty"`
	Parent ReplyParent `json:"parent,omitempty"`
}
type ReplyRoot struct {
	CID string `json:"cid,omitempty"`
	URI string `json:"uri,omitempty"`
}
type ReplyParent struct {
	CID string `json:"cid,omitempty"`
	URI string `json:"uri,omitempty"`
}
