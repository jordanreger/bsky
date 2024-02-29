package types

import "time"

type Embed struct {
	Type     string         `json:"$type"`
	External *ExternalEmbed `json:"external,omitempty"`
	Images   []*Image       `json:"images,omitempty"`
	Record   *RecordEmbed   `json:"record,omitempty"`
}

type ExternalEmbed struct {
	URI         string `json:"uri,omitempty"`
	Thumb       string `json:"thumb,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type AspectRatio struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Image struct {
	Alt         string       `json:"alt,omitempty"`
	Image       *Image       `json:"image,omitempty"`
	Type        string       `json:"$type,omitempty"`
	Ref         *Ref         `json:"ref,omitempty"`
	MimeType    string       `json:"mimeType,omitempty"`
	Size        int          `json:"size,omitempty"`
	Thumb       string       `json:"thumb,omitempty"`
	FullSize    string       `json:"fullsize,omitempty"`
	AspectRatio *AspectRatio `json:"aspectRatio,omitempty"`
}
type Ref struct {
	Link string `json:"$link,omitempty"`
}

type RecordEmbed struct {
	Type      string     `json:"$type"`
	URI       string     `json:"uri"`
	CID       string     `json:"cid"`
	Author    *Actor     `json:"author,omitempty"`
	Value     *Record    `json:"value,omitempty"`
	Labels    []string   `json:"labels,omitempty"`
	IndexedAt *time.Time `json:"indexedAt,omitempty"`
	Embeds    []Embed    `json:"embeds,omitempty"`
}
