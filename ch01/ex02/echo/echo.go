package main

/*
 * Exercise 1.2:
 * Modify the echo program to print the index and value of each of its arguments,
 * one per line.
 */
import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("#%d: %s\n", i, os.Args[i])
	}
}
