package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	url := "http://www.gopl.io"
	words, images, _ := CountWordsAndImages(url)
	fmt.Printf("%d,%d", words, images)
}
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return 0, 0, err
	}
	words, images = countWordsAndImages(doc)
	return words, images, nil

}

func countWordsAndImages(n *html.Node) (words, images int) {
	images = 0
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}

	words = 0
	//たりない. scannerでscanする.
	if n.Type == html.TextNode {
		words++
	}

	if n.FirstChild != nil {
		img, wrd := countWordsAndImages(n.FirstChild)
		images += img
		wrd += wrd
	}

	if n.NextSibling != nil {
		img, wrd := countWordsAndImages(n.NextSibling)
		images += img
		wrd += wrd
	}
	return images, words
}
