package bsky

type Feed = []FeedItem

type FeedItem struct {
	Post *Post `json:"post,omitempty"`
}
