package main

import (
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "gov.uk",
	Path:   "/government/publications/register-of-licensed-sponsors-workers",
}

// Gets a PDF link from a web-page to be parsed later on.
func getPdfLink(url string) string {
	// Get a page by an URL
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse HTML of a page
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// Recursively visit nodes in the tree
	var v string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" && len(n.Attr) > 0 {
			for _, attr := range n.Attr {
				if attr.Key == "aria-describedby" {
					v = n.Attr[len(n.Attr)-1].Val
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return v
}
