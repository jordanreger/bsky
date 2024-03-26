package bsky

import (
	"time"
)

type List struct {
	List  ListList `json:"list,omitempty"`
	Items []Actor  `json:"items,omitempty"`
}

type ListList struct {
	URI               string    `json:"uri,omitempty"`
	CID               string    `json:"cid,omitempty"`
	Name              string    `json:"name,omitempty"`
	Purpose           string    `json:"purpose,omitempty"`
	IndexedAt         time.Time `json:"indexedAt,omitempty"`
	Labels            []string  `json:"labels,omitempty"`
	Creator           Actor     `json:"creator,omitempty"`
	Description       string    `json:"description,omitempty"`
	DescriptionFacets []Facet   `json:"descriptionFacets,omitempty"`
}
