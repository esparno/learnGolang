package main

import (
	"net"
	"log"
	//"io"
	"fmt"
	"bufio"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle (conn)
	}
}
func handle (conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		fmt.Fprint(conn, ln)
		fmt.Println(ln)
	}
	defer conn.Close()
}