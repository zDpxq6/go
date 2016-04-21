// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	input, err := os.Open("./resource/Go.html")
	defer input.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	for elem, num := range visit(make(map[string]int), doc) {
		fmt.Printf("%s\t%d\n", elem, num)
	}
}

func visit(elements map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return elements
	}

	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	elements = visit(elements, n.FirstChild)
	elements = visit(elements, n.NextSibling)

	return elements
}
