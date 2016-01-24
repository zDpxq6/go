package main
/*
 * Exercise 1.1:
 * Modify the echo program to also print os.Args[0],
 * the name of the command that invoked it.
*/
import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
