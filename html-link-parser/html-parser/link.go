package parser

import (
	"golang.org/x/net/html"
)

type Link struct {
	Text string
	Href string
}

func (p *Parser) Link() []Link {
	doc, _ := html.Parse(p.r)

	var links []Link

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			link := Link{
				Text: n.FirstChild.Data,
			}
			for _, a := range n.Attr {
				if a.Key == "href" {
					link.Href = a.Val
				}
			}
			links = append(links, link)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links
}
