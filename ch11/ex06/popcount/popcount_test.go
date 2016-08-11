// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"gopl.io/ch2/popcount"
)

func PopCount24(x uint64) int {
	result := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			result++
		}
		x = x >> 1
	}
	return result
}

func PopCount25(x uint64) int {
	result := 0
	for ; x != 0; x &= x - 1 {
		result++
	}
	return result
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x12345)//67890ABCDEF)
	}
}

func BenchmarkPopCount24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount24(0x12345)//67890ABCDEF)
	}
}

func BenchmarkPopCount25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount25(0x12345)//67890ABCDEF)
	}
}


//args=0x1234
//testing: warning: no tests to run
//PASS
//BenchmarkPopCount-8  	200000000	         8.08 ns/op
//BenchmarkPopCount24-8	20000000	        76.6 ns/op
//BenchmarkPopCount25-8	200000000	         6.59 ns/op
//ok  	command-line-arguments	6.049s

//0x12345
//testing: warning: no tests to run
//PASS
//BenchmarkPopCount-8  	200000000	         8.17 ns/op
//BenchmarkPopCount24-8	20000000	        78.7 ns/op
//BenchmarkPopCount25-8	200000000	         8.93 ns/op
//ok  	command-line-arguments	6.820s
