// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg sync.WaitGroup) {
	defer fmt.Println("            echo: end")
	fmt.Println("            echo: start")
	defer fmt.Println("        wg decremented")
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn, wg sync.WaitGroup, event chan struct{}) {
	defer fmt.Println("    handleConn: end")
	fmt.Println("    handleConn: start")
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		fmt.Println("        wg incremented")
		event <- struct{}{}
		fmt.Println("        send event")
		go echo(c, input.Text(), 3*time.Second, wg)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-

func main() {
	defer fmt.Println("main: end")
	fmt.Println("main: start")
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	func(l net.Listener) {
		var wg sync.WaitGroup
		event := make(chan struct{})
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Print(err) // e.g., connection aborted
				continue
			}
			go handleConn(conn, wg, event)
			<-event
			fmt.Println("        receive event")
			wg.Wait()
			conn.(*net.TCPConn).CloseWrite()
			fmt.Println("        closed")
		}
	}(l)
}
