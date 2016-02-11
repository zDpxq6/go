// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"
)

func main() {
	output1, _ := os.Create("output1.txt")
	first := measure(output1)
	output2, _ := os.Create("output2.txt")
	second := measure(output2)
	fmt.Printf("The difference is %.2fs\n", math.Abs(first-second))
}

func measure(fp *os.File) float64 {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, fp) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	result := time.Since(start).Seconds()
	fmt.Println(strings.Join(os.Args[1:], " "))
	return result
}

func fetch(url string, ch chan<- string, fp *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(fp, resp.Body)
	fp.Close()
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	return
}

//!-
