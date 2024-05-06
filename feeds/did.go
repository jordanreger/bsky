package feeds

import (
	"encoding/json"
	"strings"
)

type Service struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

type DIDRes struct {
	ID      string    `json:"id"`
	Service []Service `json:"service"`
}

func GetWellKnownDID(endpoint string) string {
	domain := strings.ReplaceAll(endpoint, "http://", "")
	domain = strings.ReplaceAll(domain, "https://", "")

	service := Service{
		"#bsky_fg",
		"BskyFeedGenerator",
		endpoint,
	}

	did := DIDRes{
		"did:web:" + domain,
		[]Service{service},
	}

	res, err := json.Marshal(did)
	if err != nil {
		panic(err)
	}

	return string(res)
}
