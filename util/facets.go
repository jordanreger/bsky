package util

import (
	"bytes"
	"html/template"
	"regexp"
	"strings"
)

type Facet struct {
	// This field is slightly modified because Bluesky doesn't use it consistently
	Type     string         `json:"$type"`
	Index    FacetIndex     `json:"index"`
	Features []FacetFeature `json:"features"`
}

type FacetIndex struct {
	ByteEnd   int `json:"byteEnd"`
	ByteStart int `json:"byteStart"`
}

type FacetFeature struct {
	DID  string `json:"did"`
	URI  string `json:"uri"`
	Type string `json:"$type"`
}

var handleRegex = regexp.MustCompile(`[$|\W](@([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)`)

func ParseMentions(raw string) []Facet {
	var mentions []Facet

	rawBytes := []byte(raw)

	for _, m := range handleRegex.FindAllString(raw, -1) {
		did := GetDID(strings.Split(m, "@")[1])

		mentions = append(mentions,
			Facet{
				Index: FacetIndex{
					ByteStart: bytes.Index(rawBytes, []byte(m)),
					ByteEnd:   bytes.Index(rawBytes, []byte(m)) + len(m),
				},
				Features: []FacetFeature{
					{
						Type: "app.bsky.richtext.facet#mention",
						DID:  did,
					},
				},
			},
		)
	}

	return mentions
}

var urlRegex = regexp.MustCompile(`[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
var emailRegex = regexp.MustCompile(`.*@.*`)

func ParseURLs(raw string) []Facet {
	var urls []Facet

	rawBytes := []byte(raw)

	for _, u := range urlRegex.FindAllString(raw, -1) {
		if !emailRegex.MatchString(u) {
			urls = append(urls,
				Facet{
					Index: FacetIndex{
						ByteStart: bytes.Index(rawBytes, []byte(u)),
						ByteEnd:   bytes.Index(rawBytes, []byte(u)) + len(u),
					},
					Features: []FacetFeature{
						{
							Type: "app.bsky.richtext.facet#link",
							URI:  "https://" + string(u),
						},
					},
				},
			)
		}
	}
	return urls
}

func ParseFacets(text string) []Facet {
	var facets []Facet

	facets = append(facets, ParseURLs(text)...)

	facets = append(facets, ParseMentions(text)...)

	return facets
}

func FacetsToHTML(text string, facets []Facet) template.HTML {
	if len(facets) == 0 {
		return template.HTML(text)
	}
	linkFacet := "app.bsky.richtext.facet#link"
	mentionFacet := "app.bsky.richtext.facet#mention"
	offset := 0
	for _, f := range facets {
		if f.Features[0].Type == linkFacet {
			in_txt := text[f.Index.ByteStart+offset : f.Index.ByteEnd+offset]
			m_url := "<a href='" + f.Features[len(f.Features)-1].URI + "'>" + in_txt + "</a>"
			text = text[:f.Index.ByteStart+offset] + m_url + text[f.Index.ByteEnd+offset:]
			offset += len(m_url) - len(in_txt)
		} else if f.Features[0].Type == mentionFacet {
			in_txt := text[f.Index.ByteStart+offset : f.Index.ByteEnd+offset]
			m_url := "<a href='https://htmlsky.app/profile/" + f.Features[len(f.Features)-1].DID + "'>" + in_txt + "</a>"
			text = text[:f.Index.ByteStart+offset] + m_url + text[f.Index.ByteEnd+offset:]
			offset += len(m_url) - len(in_txt)
		}
	}

	return template.HTML(text)
}
