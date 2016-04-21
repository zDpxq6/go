// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	//max, minともにパッケージ外に公開していません.
	//なので, 可変長引数のmax,minの引数を空の状態で呼び出さないのは呼ぶ側の責務です.
	//max, minを呼ぶ前に, 必ず引数が空でない呼び方をするようにします.
	//
	fmt.Println(max(1, 10))
	fmt.Println(min(1, 10))
}

func sum(value int, vals ...int) int {
	total := value
	for _, val := range vals {
		total += val
	}
	return total
}

func max(val int, vals ...int) int {
	result := val
	for _, val := range vals {
		if result <= val {
			result = val
		}
	}
	return result
}
func min(val int, vals ...int) int {
	result := val
	for _, val := range vals {
		if val <= result {
			result = val
		}
	}
	return result
}
