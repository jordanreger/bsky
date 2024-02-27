package types

import "html/template"

type Actor struct {
	DID             string `json:"did"`
	Handle          string `json:"handle"`
	DisplayName     string `json:"displayName"`
	Description     string `json:"description"`
	Avatar          string `json:"avatar"`
	Banner          string `json:"banner"`
	FollowersCount  int    `json:"followersCount"`
	FollowsCount    int    `json:"followsCount"`
	PostsCount      int    `json:"postsCount"`
	Feed            []FeedItem
	DescriptionHTML template.HTML
}
