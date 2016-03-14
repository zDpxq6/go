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
	result := countElement(doc)
	for k,v := range result {
		fmt.Printf("%v:\t%v\n",k,v);
	}
}

func countElement(n *html.Node) map[string]int {
	m := make(map[string]int)
	if n.Type == html.ElementNode {
		m[n.Data] = 100
		fmt.Printf("%s:%d\n", n.Data, m[n.Data])
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Printf("    ")
		childMap := countElement(c)
		for k, _ := range childMap {
			m[k]++
		}
	}
	return m
}
