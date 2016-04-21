// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"
"strings"
	"golang.org/x/net/html"
)

func main() {
	input, err := os.Open("./resource/Go.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(input)
	defer input.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		os.Exit(1)
	}
	for _, line := range visit(nil, doc) {
		fmt.Println(line)
	}

}

func visit(text []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return text
		}
	}

	if n.Type == html.TextNode {
		e := strings.TrimSpace(n.Data)
		text = append(text, e)
	}

	if n.FirstChild != nil {
		text = visit(text, n.FirstChild)
	}

	if n.NextSibling != nil {
		text = visit(text, n.NextSibling)
	}

	return text
}
