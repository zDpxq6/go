/*
Exercise 2.5:
The expression x&(x-1) clears the right most non-zero bit of x.
Write a version of PopCount that counts bits by using this fact,
and assess its performance.
*/
package popcount

var pc [256]byte

// pc[i] is the popuration count of i;
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	result := 0
	for ; x != 0; x &= x - 1 {
		result++
	}
	return result
}
