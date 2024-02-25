package main

import (
	"html/template"
)

func parseFacets(text string, facets []Facet) template.HTML {
	offset := 0
	for _, f := range facets {
		if f.Features[0].Type == "app.bsky.richtext.facet#link" {
			in_txt := text[f.Index.ByteStart+offset : f.Index.ByteEnd+offset]
			m_url := "<a href='" + f.Features[len(f.Features)-1].URI + "'>" + in_txt + "</a>"
			text = text[:f.Index.ByteStart+offset] + m_url + text[f.Index.ByteEnd+offset:]
			offset += len(m_url) - len(in_txt)
		} else if f.Features[0].Type == "app.bsky.richtext.facet#mention" {
			in_txt := text[f.Index.ByteStart+offset : f.Index.ByteEnd+offset]
			m_url := "<a href='https://htmlsky.app/profile/" + f.Features[len(f.Features)-1].DID + "'>" + in_txt + "</a>"
			text = text[:f.Index.ByteStart+offset] + m_url + text[f.Index.ByteEnd+offset:]
			offset += len(m_url) - len(in_txt)
		}
	}

	res := template.HTML(text)

	return res
}
