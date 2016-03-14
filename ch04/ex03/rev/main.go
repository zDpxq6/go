package main

import (
	"fmt"
)

func main() {
	test1 := []int{1, 2, 3}
	fmt.Printf("before: %v\n", test1)
	reverse(&test1)
	fmt.Printf("after: %v\n", test1)
}

func reverse(s *[]int) {
	ss := *s
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
}
