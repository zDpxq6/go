// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	ExitCodeOK = iota
	ExitCodeError
)

type CharCount struct {
	outStream, errStream io.Writer
}

func (x *CharCount) Run(args []string) int {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := strings.NewReader(args[1])
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(x.errStream, "charcount: %v\n", err)
			return ExitCodeError
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Fprintf(x.outStream, "rune\tcount\r\n")
	for c, n := range counts {
		fmt.Fprintf(x.outStream, "%q\t%d\r\n", c, n)
	}
	fmt.Fprint(x.outStream, "\r\nlen\tcount\r\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(x.outStream, "%d\t%d\r\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Fprintf(x.outStream, "\n%d invalid UTF-8 characters\n", invalid)
	}
	return ExitCodeOK
}

//!-
