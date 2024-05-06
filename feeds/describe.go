package feeds

import (
	"encoding/json"
)

type DescribeFeed struct {
	URI string `json:"uri"`
}

type Describe struct {
	DID   string `json:"did"`
	Feeds []Feed `json:"feeds"`
}

func DescribeFeedGenerator(did string) string {
	describe := Describe{
		did,
		[]Feed{
			{"test"},
		},
	}

	res, err := json.Marshal(describe)
	if err != nil {
		panic(err)
	}

	return string(res)
}
