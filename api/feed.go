package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/jordanreger/htmlsky/types"
	"github.com/jordanreger/htmlsky/util"
)

type f_res struct {
	Feed   types.Feed `json:"feed"`
	Cursor string     `json:"cursor"`
}

func GetActorFeed(actor types.Actor) types.Feed {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.feed.getAuthorFeed?actor=" + actor.DID)
	if err != nil {
		fmt.Println(err)
	}

	var f_body f_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &f_body)

	feed := f_body.Feed

	var wg sync.WaitGroup

	for i := range feed {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			feed[i].Post.RKey = util.GetRKey(feed[i].Post.URI)
		}(i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			feed[i].Post.Record.HTML = util.FacetsToHTML(feed[i].Post.Record.Text, feed[i].Post.Record.Facets)
		}(i)
	}

	wg.Wait()

	return feed
}
