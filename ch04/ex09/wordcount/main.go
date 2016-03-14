// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

const(
	filename = "/Users/tsuguka/Desktop/test.txt"
)

func main() {
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	
	result := make(map[string]int)
	for scanner.Scan() {
		key := scanner.Text()
		result[key]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	file.Close()

	fmt.Println("---result---")
	for key, _ := range result {
		fmt.Printf("%v\t%d\n", key, result[key])
	}
	os.Exit(0)
}
