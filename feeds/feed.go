package feeds

import (
	"encoding/json"
)

type Feed struct {
	URI string `json:"uri"`
}

type Post struct {
	URI string `json:"post"`
}

type Skeleton struct {
	Feed []Post `json:"feed"`
}

func GetFeedSkeleton(posts []Post) string {
	skeleton := Skeleton{
		posts,
	}

	res, err := json.Marshal(skeleton)
	if err != nil {
		panic(err)
	}

	return string(res)
}
