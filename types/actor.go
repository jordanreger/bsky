package types

import "html/template"

type Actor struct {
	DID             string        `json:"did"`
	Handle          string        `json:"handle"`
	DisplayName     string        `json:"displayName,omitempty"`
	Description     string        `json:"description,omitempty"`
	Avatar          string        `json:"avatar,omitempty"`
	Banner          string        `json:"banner,omitempty"`
	FollowersCount  int           `json:"followersCount"`
	FollowsCount    int           `json:"followsCount"`
	PostsCount      int           `json:"postsCount"`
	Feed            []FeedItem    `json:"feed,omitempty"`
	DescriptionHTML template.HTML `json:"descriptionHTML,omitempty"`
}
