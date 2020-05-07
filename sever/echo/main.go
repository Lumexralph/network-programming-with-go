// Package main is an implementation of a mult-threaded
// echo server.
package main

import (
	"log"
	"net"
)

func main() {
	address := ":1300"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Fatal(err)
	}

	// bind the server to a port on the host machine
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("could not bind to port: %s\n", err)
	}

	for {
		// accept connection through handshake
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue // on no account should the server quit because of a client error
		}

		go handleEcho(conn)
	}
}

func handleEcho(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
			return
		}

		// respond to the client
		_, err = conn.Write(buf[0:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}
