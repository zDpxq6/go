package main

import (
	"fmt"
)

func main() {
	test1 := []int{1, 2, 3}

	fmt.Printf("%v", rotate(test1[:], 2))

}

func rotate(s []int, newHead int) []int {
	length := len(s)
	for _, e := range s {
		s = append(s, e)
	}
	return s[newHead : newHead+length]
}
