package bsky

import (
	"html/template"
	"jordanreger.com/bsky/util"
	"time"
)

type List struct {
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

type ListItem struct {
	URI     string `json:"uri,omitempty"`
	Subject Actor  `json:"subject,omitempty"`
}

func (list List) RKey() string {
	return util.GetRKey(list.URI)
}

func (list List) Feed() Feed {
	return GetListFeed(list.URI)
}

func (list List) DescriptionHTML() template.HTML {
	return FacetsToHTML(list.Description, list.DescriptionFacets)
}
