package bsky

import "time"

type Embed struct {
	Type     string         `json:"$type"`
	External *ExternalEmbed `json:"external"`
	Images   []*Image       `json:"images"`
	Record   *RecordEmbed   `json:"record"`
}

type ExternalEmbed struct {
	URI         string `json:"uri"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AspectRatio struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Image struct {
	Alt         string       `json:"alt"`
	Image       *Image       `json:"image"`
	Type        string       `json:"$type"`
	Ref         *Ref         `json:"ref"`
	MimeType    string       `json:"mimeType"`
	Size        int          `json:"size"`
	Thumb       string       `json:"thumb"`
	FullSize    string       `json:"fullsize"`
	AspectRatio *AspectRatio `json:"aspectRatio"`
}
type Ref struct {
	Link string `json:"$link"`
}

type RecordEmbed struct {
	Type      string     `json:"$type"`
	URI       string     `json:"uri"`
	CID       string     `json:"cid"`
	Author    *Actor     `json:"author"`
	Value     *Record    `json:"value"`
	Labels    []string   `json:"labels"`
	IndexedAt *time.Time `json:"indexedAt"`
	Embeds    []Embed    `json:"embeds"`
}
