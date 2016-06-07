package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	src := `
	<!DOCTYPE html>
	<html lang="ja">
		<head>
			<meta charset="utf8">
			<title>はじめてのHTML</title>
		</head>
		<body>
			<p>こんにちは!</p>
		</body>
	</html>`
	reader := NewReader(src)
	doc, err := html.Parse(reader)//Parse(r io.Reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}