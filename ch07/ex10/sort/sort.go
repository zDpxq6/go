package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1, 2, 3, 3, 2, 1} // unsorted
	sl := sort.IntSlice(s)
	fmt.Println(isPalindrome(sl))
}

func isPalindrome(s sort.Interface) bool {
	maxIndex := s.Len() - 1
	for i := 0; i < maxIndex/2; i++ {
		if !s.Less(i, maxIndex - i) && !s.Less(maxIndex - i, i) {
			continue
		}
		return false
	}
	return true
}
