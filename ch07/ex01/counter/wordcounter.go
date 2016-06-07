package main

import (
	"bufio"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	var counter int
	for scanner.Scan() {
		scanner.Text()
		counter++
	}
	*c = WordCounter(counter)
	return counter, nil
}
