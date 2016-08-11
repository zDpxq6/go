package main

import (
	"fmt"
	"time"
)

func main() {

	// Create pipeline
	start := make(chan struct{})
	var in *chan struct{}
	var last *chan struct{}
	for i := 0; ; i++ {
		if i == 0 {
			in = &start
		} else {
			in = last
		}

		out := make(chan struct{})
		pass(*in, out)
		last = &out
	}

	timer := time.Now()
	start <- struct{}{}
	<-(*last)
	fmt.Printf("Run: %s\n", time.Since(timer))
}

func pass(in <-chan struct{}, out chan<- struct{}) {
	go func() {
		out <- <-in
	}()
}
