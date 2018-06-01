package adventures

import (
	"encoding/json"
	"io"
)

// Adventures represent the adventures.
type Adventures = map[string]Adventure

// Adventure is the main object
type Adventure struct {
	Title   string   `json:"title,omitempty"`
	Story   []string `json:"story,omitempty"`
	Options []struct {
		Text string `json:"text,omitempty"`
		Arc  string `json:"arc,omitempty"`
	} `json:"options,omitempty"`
}

// New gives back all the Adventures
func New(r io.Reader) *Adventures {
	a := &Adventures{}
	json.NewDecoder(r).Decode(a)
	return a
}

// Find finds the next Adventure
func Find(a Adventures, k string) Adventure {
	return a[k]
}
