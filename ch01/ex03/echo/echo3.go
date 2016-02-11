package main

/*
 * Exercise 1.2:
 * Modify the echo program to print the index and value of each of its arguments,
 * one per line.
 */
import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	fmt.Println("by +=")
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	result := time.Since(start).Seconds()
	fmt.Println(result)

	fmt.Println("by Join")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	result = time.Since(start).Seconds()
	fmt.Println(result)
}
