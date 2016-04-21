// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Join("", "test", "test", "test", "test"))
}

func Join(seq string, a ...string) string {
	if a == nil {
		return ""
	}
	var result string
	for _, val := range a {
		result += val + seq
	}

	return result[:len(result)-len(seq)]
}
