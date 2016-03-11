// © 2016 zDpxq6

package main

import (
	"bytes"
	"fmt"
)

func main() {
	//	for i, r := range "Hello, 世界" {
	//		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	//	}
	fmt.Println(comma2("あ"))
	fmt.Println(comma2("あい"))
	fmt.Println(comma2("あいう"))
	fmt.Println(comma2("うあいう"))
	fmt.Println(comma2("いうあいう"))
	fmt.Println(comma2("あいうあいう"))
	fmt.Println(comma2("うあいうあいう"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	length := 0
	for range s {
		length++
	}
	if length <= 3 {
		return s
	}
	skip := length % 3
	flag := skip == 0//最初にカンマが必要な場合true
	var buf bytes.Buffer
	counter := 0
	for _, r := range s {
		if skip != 0 {
			buf.WriteRune(r)
			skip--
			continue
		}
		if counter == 0 {
			if flag {
				flag = false
			} else {
				buf.WriteRune(',')
			}
		} else if counter == 2 {
			counter = -1
		}
		counter++
		buf.WriteRune(r)
	}

	return buf.String()
}
