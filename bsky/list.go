package bsky

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetList(at_uri string) List {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.graph.getList?list=" + at_uri)
	if err != nil {
		fmt.Println(err)
	}

	var list List
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &list)

	return list
}
