package urlshort

import (
	"io"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type PathToUrl struct {
	Path string `yaml:"path,omitempty"`
	URL  string `yaml:"url,omitempty"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for k, v := range pathsToUrls {
			if strings.HasPrefix(r.URL.Path, k) {
				http.Redirect(w, r, v, http.StatusMovedPermanently)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(r io.Reader, fallback http.Handler) (http.HandlerFunc, error) {
	p := []PathToUrl{}
	pm := make(map[string]string)

	err := yaml.NewDecoder(r).Decode(&p)
	if err != nil {
		return nil, err
	}

	for _, v := range p {
		pm[v.Path] = v.URL
	}

	return MapHandler(pm, fallback), nil
}
