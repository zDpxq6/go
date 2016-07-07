// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.Int("port", 8000, "port")
	timezone := flag.String("TZ", "Asia/Tokyo", "time zone")
	flag.Parse()
	fmt.Println(*port)
	fmt.Println(*timezone)
	loc, err := time.LoadLocation(*timezone)
	if err != nil {
		fmt.Println("faild to load a location")
	}
	getTime(*port, loc)
}

func getTime(port int, loc *time.Location) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, loc) // handle connections concurrently
	}

}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format(time.RFC1123)+"\n")
		if err != nil {
			fmt.Println("faild to write time")
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
