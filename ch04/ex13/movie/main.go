// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	requestTemplate = "http://www.omdbapi.com/?t=%s&r=json"
)

//!+
func main() {
	name := os.Args[1]
	req := fmt.Sprintf(requestTemplate, name)
	res, err := http.Get(req)
	if err != nil {
		fmt.Printf("Getting movie info failed: %s", err)
		os.Exit(1)
	}
	
	defer res.Body.Close()
	info, err := ioutil.ReadAll(res.Body)
	var posterURL struct{ Poster string }
	if err := json.Unmarshal(info, &posterURL); err != nil {
		fmt.Printf("JSON unmarshaling failed: %s", err)
		os.Exit(1)
	}

	pos, err := http.Get(posterURL.Poster)
	if err != nil {
		fmt.Printf("Getting poster failed: %s", err)
		os.Exit(1)
	}
	defer pos.Body.Close()
	poster, err := ioutil.ReadAll(pos.Body)

	var outFile *os.File
	if outFile, err = os.Create(name + ".jpg"); err != nil {
		fmt.Printf("Creating poster failed: %s", err)
		os.Exit(1)
	}

	img, err := jpeg.Decode(bytes.NewReader(poster))
	if err != nil {
		fmt.Printf("Decording failed: %s", err)
		os.Exit(1)
	}

	option := &jpeg.Options{Quality: 100}
	if err = jpeg.Encode(outFile, img, option); err != nil {
		fmt.Printf("Encording failed: %s", err)
		os.Exit(1)
	}

	defer outFile.Close()
	os.Exit(0)
}
