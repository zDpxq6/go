// © 2016 zDpxq6

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", isAnagram("a", "ab"))
	fmt.Printf("%v\n", isAnagram("aa", "ab"))
	fmt.Printf("%v\n", isAnagram("aa", "aa"))
	fmt.Printf("%v\n", isAnagram("ab", "ba"))
	fmt.Printf("%v\n", isAnagram("世界", "界世"))
}

func isAnagram(s, t string) bool {
	if s == "" || t == "" {
		return true
	}
	if len(s) != len(t) {
		return false
	}
	return sortString(s) == sortString(t)
}

func sortString(s string) string {
	return intSliceToString(stringToSortedIntSlice(s))
}

func stringToSortedIntSlice(s string) []int {
	size := len(s)
	ints := make([]int, size, size*2)
	runes := []rune(s)
	for i, r := range runes {
		ints[i] = int(r)
	}
	result := sort.IntSlice(ints)
	result.Sort()
	return result
}

func intSliceToString(is []int) string {
	size := len(is)
	runes := make([]rune, size, size*2)
	for i, e := range is {
		runes[i] = rune(e)
	}
	return string(runes)
}
