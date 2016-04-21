// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"strings"
)

const (
	prefix    = "$"
	remaining = 2
)

func main() {
	target := "$foo"
	f := func(s string) string {
		return s
	}
	fmt.Println(expand(target, f))
}
func extractSubstring(original, starter string) (string, bool) {
	if !strings.Contains(original, starter) {
		return "", false
	}
	splitted := strings.SplitN(original, starter, remaining)
	return splitted[1], true
}
func expand(s string, f func(string) string) string {
	substring, ok := extractSubstring(s, prefix)
	if !ok {
		return ""
	}
	return f(substring)
}
