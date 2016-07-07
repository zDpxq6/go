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
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	defer fmt.Println("    handleConn: end")
	fmt.Println("    handleConn: start")
	readevent := make(chan struct{})

	go func(c net.Conn, readevent chan struct{}) {
		defer fmt.Println("        monitor input: end")
		fmt.Println("        monitor input: start")
		for {
			c.Read(make([]byte, 1))
			readevent <- struct{}{}
		}
	}(c, readevent)

	input := bufio.NewScanner(c)
	for input.Scan() {
		ticker := time.NewTicker(10 * time.Second)
		select {
		case <-ticker.C:
			c.Close()
		case <-readevent:
			go echo(c, input.Text(), 1*time.Second)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	defer fmt.Println("main: end")
	fmt.Println("main: start")
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main: loop start")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
	fmt.Println("main: loop end")
}
