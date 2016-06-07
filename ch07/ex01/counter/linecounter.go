package main

import (
	"bufio"
	"strings"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	var counter int
	for scanner.Scan() {
		scanner.Text()
		counter++
	}
	*c = LineCounter(counter)
	return counter, nil
}
