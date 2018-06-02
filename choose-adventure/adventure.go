package adventures

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
)

var (
	// ErrInvalidStructure happends when the the caller passes an invalid structure.
	ErrInvalidStructure = errors.New("was expected to get an adventure structure")
	// ErrInvalidTemplatePath happends when the the caller passes a path and we can't find the file.
	ErrInvalidTemplatePath = errors.New("couldn't open the file")
)

var defaultQuery = "adventure"

// Game is the core object which the user will interact with.
type Game struct {
	a    Adventures
	tmpl *template.Template
}

// Adventure represent's the basic struct for an Adventure.
type Adventure struct {
	Title   string   `json:"title,omitempty"`
	Story   []string `json:"story,omitempty"`
	Options []struct {
		Text string `json:"text,omitempty"`
		Arc  string `json:"arc,omitempty"`
	} `json:"options,omitempty"`
}

// Adventures represent a map of string -> Adventure.
type Adventures = map[string]Adventure

// New takes a reader in and a TemplatePath then return a pointer to aGame
func New(r io.Reader, tempPath string) (*Game, error) {
	var err error
	a := Adventures{}

	tmpl, err := template.ParseFiles(tempPath)
	if err != nil {
		return nil, ErrInvalidTemplatePath
	}

	err = json.NewDecoder(r).Decode(&a)
	if err != nil {
		return nil, ErrInvalidStructure
	}

	return &Game{
		a:    a,
		tmpl: tmpl,
	}, nil
}

// WriteByKey takes a io.Writer and a key it then tries to find the correct adventure and write it to the template.
func (g *Game) WriteByKey(w io.Writer, k string) {
	g.tmpl.Execute(w, g.a[k])
}
