package bsky

import (
	"time"

	"git.sr.ht/~jordanreger/bsky/util"
)

type Embed struct {
	Type     string         `json:"$type"`
	External *ExternalEmbed `json:"external,omitempty"`
	Images   []*Image       `json:"images,omitempty"`
	Media    *MediaEmbed    `json:"media,omitempty"`
	Record   *RecordEmbed   `json:"record,omitempty"`
}

type ExternalEmbed struct {
	Type        string `json:"$type"`
	URI         string `json:"uri"`
	Thumb       string `json:"thumb,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type MediaEmbed struct {
	Type   string   `json:"$type"`
	Images []*Image `json:"images"`
}

type AspectRatio struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Image struct {
	Alt         string       `json:"alt,omitempty"`
	Image       *Image       `json:"image,omitempty"`
	Type        string       `json:"$type"`
	Ref         *Ref         `json:"ref,omitempty"`
	MimeType    string       `json:"mimeType,omitempty"`
	Size        int          `json:"size,omitempty"`
	Thumb       string       `json:"thumb,omitempty"`
	FullSize    string       `json:"fullsize,omitempty"`
	AspectRatio *AspectRatio `json:"aspectRatio,omitempty"`
}
type Ref struct {
	Link string `json:"$link"`
}

type RecordEmbed struct {
	RecordEmbedRecord `json:",omitempty"`
	Type              string             `json:"$type,omitempty"`
	Record            *RecordEmbedRecord `json:"record,omitempty"`
}

type RecordEmbedRecord struct {
	URI       string     `json:"uri,omitempty"`
	CID       string     `json:"cid,omitempty"`
	Author    *Actor     `json:"author,omitempty"`
	Value     *Record    `json:"value,omitempty"`
	Labels    []string   `json:"labels,omitempty"`
	IndexedAt *time.Time `json:"indexedAt,omitempty"`
	Embeds    []*Embed   `json:"embeds,omitempty"`
}

func (record RecordEmbedRecord) RKey() string {
	return util.GetRKey(record.URI)
}
