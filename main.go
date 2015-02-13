package main

import (
	"net"
	"log"
	"io"
)

func handleConnection(c io.ReadWriteCloser) {
	io.Copy(c, c)
}

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
