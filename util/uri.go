package util

import "strings"

func GetPostURI(did string, rkey string) string {
	return "at://" + did + "/app.bsky.feed.post/" + rkey
}

func GetListURI(did string, rkey string) string {
	return "at://" + did + "/app.bsky.graph.list/" + rkey
}

func GetRKey(uri string) string {
	split := strings.Split(uri, "/")
	return split[len(split)-1]
}
