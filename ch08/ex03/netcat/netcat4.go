// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
	"fmt"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000") //localhost:8080への接続を生成する
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{}) //チャネルの作成
	go transcript(os.Stdout, conn, done)
	mustCopy(conn, os.Stdin)
	conn.(*net.TCPConn).CloseRead()
	fmt.Println("closed")
	<-done // wait for background goroutine to finish
	fmt.Println("end")
}

func transcript(dst io.Writer, src io.Reader, done chan struct{}) {
	io.Copy(dst, src)  //connから読んで, 標準出力に書き出す
	done <- struct{}{} // 終了イベント送信
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
