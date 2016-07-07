package main

import (
	"net"
	"os"
	"io"
)

func main() {
	conn1, _ := net.Dial("tcp","localhost:8010")
	conn2, _ := net.Dial("tcp","localhost:8020")
//	conn3, _ := net.Dial("tcp","localhost:8030")
	defer conn1.Close()
	defer conn2.Close()
//	defer conn3.Close()
	io.Copy(os.Stdout, conn1)
	io.Copy(os.Stdout, conn2)
//	io.Copy(os.Stdout, conn3)
//	for _, v := range os.Args[1:] {
//		i := strings.Index(v, "=")
//		if i < 0 {
//			fmt.Println("illegal format")
//			continue
//		}
//		val := strings.Split(v, "=")
//		connect(val[0], val[1])
//
//	}
}

//func connect(title, address string) {
//	fmt.Println("connect")
//	conn, err := net.Dial("tcp", address)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//	r := bufio.NewReader(conn)
//	for {
//		line, err := r.ReadString('\n')
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(line)
//	}
//}
