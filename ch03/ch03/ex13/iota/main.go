// © 2016 zDpxq6

package main

import (
	"fmt"
)

const (
	base = 1000
	KB   = base
	MB   = KB * base
	GB   = MB * base
	TB   = GB * base
	PB   = TB * base
	EB   = PB * base
	ZB   = EB * base
	YB   = ZB * base
)

func main() {
	fmt.Printf("%v\n", KB)
	fmt.Printf("%v\n", MB)
	fmt.Printf("%v\n", GB)
	fmt.Printf("%v\n", TB)
	fmt.Printf("%v\n", PB)
	fmt.Printf("%v\n", EB)
	//	fmt.Printf("%v\n", float64(ZB))
	//	fmt.Printf("%v\n", float64(YB))
}
