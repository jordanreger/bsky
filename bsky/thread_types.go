package bsky

import (
	"html/template"
	"time"

	"github.com/jordanreger/htmlsky/util"
)

type Thread struct {
	Type    string    `json:"$type"`
	Post    *Post     `json:"post,omitempty"`
	Replies []*Thread `json:"replies,omitempty"`
}

type Post struct {
	URI         string     `json:"uri"`
	CID         string     `json:"cid"`
	Author      *Actor     `json:"author"`
	Record      *Record    `json:"record,omitempty"`
	Embed       *Embed     `json:"embed,omitempty"`
	ReplyCount  int        `json:"replyCount"`
	RepostCount int        `json:"repostCount"`
	LikeCount   int        `json:"likeCount"`
	IndexedAt   *time.Time `json:"indexedAt,omitempty"`
	Labels      []string   `json:"labels,omitempty"`
}

func (post Post) RKey() string {
	return util.GetRKey(post.URI)
}

type Record struct {
	Text      string     `json:"text"`
	Type      string     `json:"$type"`
	Langs     []string   `json:"langs,omitempty"`
	Embed     *Embed     `json:"embed,omitempty"`
	Facets    []*Facet   `json:"facets,omitempty"`
	Reply     *Reply     `json:"reply,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

func (record Record) HTML() template.HTML {
	return FacetsToHTML(record.Text, record.Facets)
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
