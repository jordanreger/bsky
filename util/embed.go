package util

type Embed struct {
	Type string `json:"$type"`
	ExternalEmbed
	ImageEmbed
	RecordEmbed
}

type ExternalEmbed struct {
	External External `json:"external"`
}
type External struct {
	URI         string `json:"uri"`
	Thumb       Thumb  `json:"thumb"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Thumb struct {
	Type     string `json:"$type"`
	Ref      Ref    `json:"ref"`
	MimeType string `json:"mimeType"`
	Size     int    `json:"size"`
}
type Ref struct {
	Link string `json:"$link"`
}

type ImageEmbed struct {
	Images []Image `json:"images"`
}
type Image struct {
	Thumb    string `json:"thumb"`
	FullSize string `json:"fullsize"`
	Alt      string `json:"alt"`
}

type RecordEmbed struct {
}
