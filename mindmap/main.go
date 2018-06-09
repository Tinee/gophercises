package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	link "github.com/Tinee/gophercises/html-link-parser"
)

/*
   1. GET the webpage
   2. parse all the links on the page
   3. build proper urls with our links
   4. filter out any links w/ a diff domain
   5. Find all pages (BFS)
   6. print out XML
*/

type filterFunc func(*url.URL) bool

type Scraper struct {
	root    url.URL
	filters []filterFunc
	// Seen is a map that holds the normalized URLs that points if we have been there or not { "localhost:8000" -> true }
	Seen map[string]bool
}

func main() {
	u := flag.String("url", "http://gophercises.com/", "")
	url, err := url.Parse(*u)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	s := Scraper{
		root: *url,
	}

	s.Scrape()
}

func (s *Scraper) Scrape() error {
	res, err := http.Get(s.root.String())
	if err != nil {
		return err
	}

	links, err := link.Parse(res.Body)
	if err != nil {
		return err
	}

	s.parseLinks(links)

	return nil
}

func (s *Scraper) parseLinks(l []link.Link) {
	var temp []string
	for _, v := range l {
		if strings.HasPrefix(v.Href, "/") {
			temp = append(temp, v.Href)
		} else if strings.HasPrefix(v.Href, "http") {
			temp = append(temp, v.Href)
		}
	}
	var temp2 []url.URL
	for _, u := range temp {
		purl, err := url.Parse(u)
		if err != nil {
			continue
		}

		if purl.Host == "" || purl.Host == s.root.Host {
			temp2 = append(temp2, *purl)
		}
		fmt.Println("ehsa")
	}

}

func (s *Scraper) addFilter(f filterFunc) {
	s.filters = append(s.filters, f)
}
