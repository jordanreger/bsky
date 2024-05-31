package bsky_test

import (
	"fmt"
	"testing"

	"git.sr.ht/~jordanreger/bsky"
)

func TestFacets_Empty(t *testing.T) {
	facets := bsky.ParseFacets("No facets here")
	t.Log(facets)
	if len(facets) != 0 {
		t.Fatal("Had " + fmt.Sprint(len(facets)) + " facets, wanted 0")
	}
}

func TestFacets_URLs(t *testing.T) {
	facets := bsky.ParseFacets("example.com there's two urls here https://example.com")
	for _, x := range facets {
		t.Log("URL: " + x.Features[0].URI)
	}
	if len(facets) != 2 {
		t.Fatal("Had " + fmt.Sprint(len(facets)) + " facets, wanted 2")
	}
}

func TestFacets_Mentions(t *testing.T) {
	facets := bsky.ParseFacets("@jordanreger.com @jordanr.com @jordanreger.com")
	for _, x := range facets {
		t.Log("Mention: " + x.Features[0].DID)
	}
	if len(facets) != 2 {
		t.Fatal("Had " + fmt.Sprint(len(facets)) + " facets, wanted 2")
	}
}

func TestFacets_Tags(t *testing.T) {
	facets := bsky.ParseFacets("#test test #test")
	for _, x := range facets {
		t.Log("Tag: " + x.Features[0].Tag)
	}
	if len(facets) != 2 {
		t.Fatal("Had " + fmt.Sprint(len(facets)) + " facets, wanted 2")
	}
}
