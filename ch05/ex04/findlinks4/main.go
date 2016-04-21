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
			links = append(links, extractAttributeValue(n.Attr, "src"))
		} else if n.Data == "link" {
			links = append(links, extractAttributeValue(n.Attr, "src"))
		} else if n.Data == "a" {
			links = append(links, extractAttributeValue(n.Attr, "href"))
		} else if n.Data == "script" {
			links = append(links, extractAttributeValue(n.Attr, "href"))
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

func extractAttributeValue(attributes []html.Attribute, name string) string {
	for _, attr := range attributes {
		if name == attr.Key {
			return attr.Val
		}
	}
	return ""
}
