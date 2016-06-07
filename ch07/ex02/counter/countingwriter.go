package main

import (
	"io"
)

type CountWriter struct {
	writer io.Writer
	count  int64
}

func CountingWriter(input io.Writer) (io.Writer, *int64) {
	result := &CountWriter {
		
		writer: input,
		count:0,
	}
	return result, &result.count
}

func (this *CountWriter)Write(p []byte)( int, error){
	i,e := this.Write(p)
	this.count += int64(i)
	return i, e
}
