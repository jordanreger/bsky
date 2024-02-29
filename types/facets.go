package types

type Facet struct {
	Type     string         `json:"$type"`
	Index    FacetIndex     `json:"index"`
	Features []FacetFeature `json:"features"`
}

type FacetIndex struct {
	ByteEnd   int `json:"byteEnd"`
	ByteStart int `json:"byteStart"`
}

type FacetFeature struct {
	DID  string `json:"did,omitempty"`
	URI  string `json:"uri,omitempty"`
	Type string `json:"$type"`
}
