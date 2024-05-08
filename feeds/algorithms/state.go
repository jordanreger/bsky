package algorithms

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	bsky "jordanreger.com/bsky/api"
	"jordanreger.com/bsky/feeds"
	"net/http"
	"slices"
)

type state_res struct {
	Posts []bsky.Post `json:"posts"`
}

type res struct {
	Posts []feeds.Post `json:"posts"`
}

func State(state string) ([]feeds.Post, error) {
	var posts []feeds.Post

	states := []string{"al", "ak", "az", "ar", "as", "ca", "co", "ct", "de", "dc", "fl", "ga", "gu", "hi", "id", "il", "in", "ia", "ks", "ky", "la", "me", "md", "ma", "mi", "mn", "ms", "mo", "mt", "ne", "nv", "nh", "nj", "nm", "ny", "nc", "nd", "mp", "oh", "ok", "or", "pa", "pr", "ri", "sc", "sd", "tn", "tx", "tt", "ut", "vt", "va", "vi", "wa", "wv", "wi", "wy"}

	if !slices.Contains(states, state) {
		return nil, errors.New("Not a state")
	}

	res, err := http.Get("https://public.api.bsky.app/xrpc/app.bsky.feed.searchPosts?q=%23" + state + "wx&limit=50&sort=latest")
	if err != nil {
		fmt.Println(err)
	}

	var s state_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &s)

	for _, state := range s.Posts {
		posts = append(posts, feeds.Post{URI: state.URI})
	}

	return posts, nil
}
