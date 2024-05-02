package bsky

import (
	"html/template"
	"time"
)

type Actor struct {
	DID            string        `json:"did"`
	Handle         string        `json:"handle"`
	DisplayName    string        `json:"displayName,omitempty"`
	Description    string        `json:"description,omitempty"`
	Avatar         string        `json:"avatar,omitempty"`
	Banner         string        `json:"banner,omitempty"`
	FollowsCount   int           `json:"followsCount,omitempty"`
	FollowersCount int           `json:"followersCount,omitempty"`
	PostsCount     int           `json:"postsCount,omitempty"`
	IndexedAt      *time.Time    `json:"indexedAt,omitempty"`
	Labels         []interface{} `json:"labels,omitempty"`
}

func (actor Actor) Feed() Feed {
	return GetActorFeed(actor)
}

func (actor Actor) DescriptionHTML() template.HTML {
	descFacets := ParseFacets(actor.Description)
	return FacetsToHTML(actor.Description, descFacets)
}
