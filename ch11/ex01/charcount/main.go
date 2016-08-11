// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"os"
)

func main() {
	c := &CharCount{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(c.Run(os.Args))
}

//!-
