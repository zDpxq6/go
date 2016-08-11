// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+test
package main

import (
	//	"os"
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	x := &CharCount{outStream: outStream, errStream: errStream}
	args := strings.Split("./charcount test", " ")
	x.Run(args)
	expect := `rune	count
't'	2
'e'	1
's'	1

len	count
1	4
2	0
3	0
4	0
`
	if outStream.String() != expect {
		t.Errorf("Run(%q) is:\r\n%v, expected is:\r\n%s", args, outStream.String(),expect)
	}
}
