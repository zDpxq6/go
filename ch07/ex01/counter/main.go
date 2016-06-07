package main

import (
	"fmt"
)

func main() {
	var words WordCounter
	c, _ := words.Write([]byte("this methods returns 4"))
	fmt.Println(c)
	
	var lines LineCounter
	c, _ = lines.Write([]byte("this methods \r\nreturns 2"))
	fmt.Println(c)
}
