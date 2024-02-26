package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type did_body struct {
	DID string `json:"did"`
}

func getDID(handle string) string {
	// if handle is DID just return it
	if strings.Contains(handle, "did:") {
		return handle
	}
	// otherwise resolve it
	res, err := http.Get("https://api.bsky.app/xrpc/com.atproto.identity.resolveHandle?handle=" + handle)
	if err != nil {
		fmt.Println(err)
	}
	var db did_body
	b, _ := io.ReadAll(res.Body)
	json.Unmarshal(b, &db)

	return db.DID
}
