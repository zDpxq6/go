// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.

// Findlinks1 prints the links in an HTML document read from standard input.
package findlinks

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	input, err := os.Open("./resource/Go.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(input)
	defer input.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

// visit appends to links each link found in n and returns the result.
func visit(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Printf("%s\n", a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
