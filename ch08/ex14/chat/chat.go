// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
)

//!+broadcaster
type client struct {
	channel chan string // an outgoing message channel
	name    string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func participants(clientSet map[client]bool) string {
	var buf bytes.Buffer
	for cli, _ := range clientSet {
		if buf.String() == "" {
			buf.WriteString("participants: ")
		} else {
			buf.WriteString(", ")
		}
		buf.WriteString(cli.name)
	}
	return buf.String()
}

func broadcaster() {
	clientSet := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clientSet {
				cli.channel <- msg
			}

		case cli := <-entering:
			clientSet[cli] = true
			cli.channel <- participants(clientSet)
		case cli := <-leaving:
			delete(clientSet, cli)
			close(cli.channel)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	cli := client{}
	cli.channel = make(chan string) // outgoing client messages
	go clientWriter(conn, cli.channel)

	who := conn.RemoteAddr().String()
	input:=bufio.NewScanner(conn)
	if input.Scan() {
		who = input.Text()
	}
	cli.channel <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
