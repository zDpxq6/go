// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	filename = "/Users/tsuguka/Desktop/test.txt"
)

func main() {
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	counts := make(map[rune]int)        // counts of Unicode characters
	controlCounts := make(map[rune]int) // counts of Unicode characters
	digitCounts := make(map[rune]int)   // counts of Unicode characters
	graphicCounts := make(map[rune]int) // counts of Unicode characters
	letterCounts := make(map[rune]int)  // counts of Unicode characters
	lowerCounts := make(map[rune]int)   // counts of Unicode characters
	numberCounts := make(map[rune]int)  // counts of Unicode characters
	printCounts := make(map[rune]int)   // counts of Unicode characters
	markCounts := make(map[rune]int)    // counts of Unicode characters
	punctCounts := make(map[rune]int)   // counts of Unicode characters
	spaceCounts := make(map[rune]int)   // counts of Unicode characters
	symbolCounts := make(map[rune]int)  // counts of Unicode characters
	titleCounts := make(map[rune]int)   // counts of Unicode characters
	upperCounts := make(map[rune]int)   // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int     // count of lengths of UTF-8 encodings
	invalid := 0                        // count of invalid UTF-8 characters

	in := bufio.NewReader(file)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		if unicode.IsControl(r) {
			controlCounts[r]++
		}
		if unicode.IsDigit(r) {
			digitCounts[r]++
		}
		if unicode.IsGraphic(r) {
			graphicCounts[r]++
		}
		if unicode.IsLetter(r) {
			letterCounts[r]++
		}
		if unicode.IsLower(r) {
			lowerCounts[r]++
		}
		if unicode.IsMark(r) {
			markCounts[r]++
		}
		if unicode.IsNumber(r) {
			numberCounts[r]++
		}
		if unicode.IsPrint(r) {
			printCounts[r]++
		}
		if unicode.IsPunct(r) {
			punctCounts[r]++
		}
		if unicode.IsSpace(r) {
			spaceCounts[r]++
		}
		if unicode.IsSymbol(r) {
			symbolCounts[r]++
		}
		if unicode.IsTitle(r) {
			titleCounts[r]++
		}
		if unicode.IsUpper(r) {
			upperCounts[r]++
		}

		utflen[n]++
	}
	file.Close()
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if 1 < len(controlCounts) {
		fmt.Printf("control\tcount\n")
		for c, n := range controlCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(digitCounts) {
		fmt.Printf("digit\tcount\n")
		for c, n := range digitCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(graphicCounts) {
		fmt.Printf("graphic\tcount\n")
		for c, n := range graphicCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(letterCounts) {
		fmt.Printf("letter\tcount\n")
		for c, n := range letterCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(lowerCounts) {
		fmt.Printf("lower\tcount\n")
		for c, n := range lowerCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(markCounts) {
		fmt.Printf("mark\tcount\n")
		for c, n := range markCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(numberCounts) {
		fmt.Printf("number\tcount\n")
		for c, n := range numberCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(printCounts) {
		fmt.Printf("print\tcount\n")
		for c, n := range printCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(punctCounts) {
		fmt.Printf("punct\tcount\n")
		for c, n := range punctCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(spaceCounts) {
		fmt.Printf("space\tcount\n")
		for c, n := range spaceCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(symbolCounts) {
		fmt.Printf("symbol\tcount\n")
		for c, n := range symbolCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(titleCounts) {
		fmt.Printf("title\tcount\n")
		for c, n := range titleCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if 1 < len(upperCounts) {
		fmt.Printf("upper\tcount\n")
		for c, n := range upperCounts {
			fmt.Printf("%q\t%d\n", c, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	os.Exit(0)
}
