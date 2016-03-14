package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

const separator = " "

var input = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	if input.Scan() {
		s = input.Text()
	}
	strings := strings.Split(s, " ")
	if 2 < len(strings) {
		return
	}
	if len(strings) == 1 {
		result := sha256.Sum256([]byte(strings[0]))
		fmt.Printf("%x", result)
	} else if strings[1] == "-sha512" {
		result := sha512.Sum512([]byte(strings[0]))
		fmt.Printf("%x", result)
	} else if strings[1] == "-sha384" {
		result := sha512.Sum384([]byte(strings[0]))
		fmt.Printf("%x", result)
	}
	return
}
