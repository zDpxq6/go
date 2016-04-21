package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	input, err := os.Open("./resource/Wikipedia.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			for _, a := range n.Attr {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
