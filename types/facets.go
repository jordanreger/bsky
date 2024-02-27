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
	DID  string `json:"did"`
	URI  string `json:"uri"`
	Type string `json:"$type"`
}
