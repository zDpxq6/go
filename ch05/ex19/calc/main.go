// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(panicer())
}

func panicer() string {
	defer func() string {
		err := recover()
		if err == nil {
			return "AAA"
		}
		return "BBB" //ここにきてるが...
	}()
	panic("Panic!!")
}
