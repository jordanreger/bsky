package bsky

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type li_res struct {
	List  List       `json:"list,omitempty"`
	Items []ListItem `json:"items,omitempty"`
}

func GetList(at_uri string) List {
	res, err := http.Get("https://api.bsky.app/xrpc/app.bsky.graph.getList?list=" + at_uri)
	if err != nil {
		fmt.Println(err)
	}

	var l_body li_res
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &l_body)

	list := l_body.List
	list.RKey()

	return list
}
