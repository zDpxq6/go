package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan struct{})
	pong := make(chan struct{})
//	counter := make(chan struct{})
	i := 0
	go func() {
		for {
			<-ping
//			counter <- struct{}{}
			pong <- struct{}{}
			fmt.Println("ping")
		}
	}()
	go func() {
		for {
			<-pong
//			counter <- struct{}{}
			ping <- struct{}{}
			fmt.Println("pong")
		}
	}()
	ping <- struct{}{}
//	go func() {
//		<-counter
//		i++
//	}()
	go func() {
		time.Sleep(time.Second)
		close(ping)
		close(pong)
//		close(counter)
		fmt.Println(i)
	}()
}
