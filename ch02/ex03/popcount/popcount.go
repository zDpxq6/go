/*
Exercise 2.3:
Rewrite PopCount to use a loop instead of a single expression.
Compare the performance of the two versions.
(Section 11.4 shows how to compare the performance of different implementations systematically.)
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
	var mask uint64 = 1
	for ; mask != 0; mask = mask << 1 {
		if mask&x != 0 {
			result++
		}
	}
	return result
}
