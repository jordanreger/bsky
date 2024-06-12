package bsky

import (
	"net/url"
)

type Agent struct {
	Service url.URL

	API API
}

type API struct {
	Com Com
}
