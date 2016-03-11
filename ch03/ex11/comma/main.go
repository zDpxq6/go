// © 2016 zDpxq6

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	separator rune = ','
)

func main() {
	//	for i, r := range "Hello, 世界" {
	//		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	//	}
	fmt.Println(comma(3.1))
	fmt.Println(comma(3.14))
	fmt.Println(comma(3.141))
	fmt.Println(comma(3.1415))
	fmt.Println(comma(3.14159))
	fmt.Println(comma(3.141592))
	fmt.Println(comma(3.1415926))
	fmt.Println(comma(3.14159265))
	fmt.Println(comma(3.141592659))
	fmt.Println(comma(3.1415926593))
	fmt.Println(comma(3.14159265935))
	fmt.Println("------------------")
	fmt.Println(comma(1.0))
	fmt.Println(comma(12.0))
	fmt.Println(comma(123.0))
	fmt.Println(comma(3123.0))
	fmt.Println(comma(23123.0))
	fmt.Println(comma(123123.0))
}

func comma(v float64) string {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	index := strings.Index(s, ".")
	integer := s[0:index]
	float := s[index+1:]
	return separateInteger(integer) + "." + separateFloat(float)
}

func separateInteger(s string) string {
	fmt.Printf("\t引数: %v\n", s)
	if strings.HasPrefix(s, "-") {
		return separateInteger(s[1:])
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:n-3] + " " + separateInteger(s[n-3:])
}

func separateFloat(s string) string {
	//	fmt.Printf("\t引数: %v\n", s)
	n := len(s)
	if n <= 3 {
		return s
	}
	if len(s[3:]) == 1 {
		return s
	}
	return s[:3] + " " + separateFloat(s[3:])
}
