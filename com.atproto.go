package bsky

import "net/url"

type Com struct {
	Atproto Atproto
}

type Atproto struct{}

func (a Atproto) URL() *url.URL {
	u, _ := url.Parse("https://atproto.com")
	return u
}
