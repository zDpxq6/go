package main

import (
	"fmt"
)

func main() {
	test1 := []string{"aa", "bb", "cc", "cc", "dd", "dd", "dd", "ee"}

	fmt.Printf("%v", deduplicate(test1[:]))

}

func deduplicate(s []string) []string{
	edge := 0;
	for i, e := range s {
		if i == 0 {
			edge++
			continue
		}
		if e == s[edge - 1] {
			continue
		} 
		s[edge] = e
		edge++
	}
	return s[:edge]
}
