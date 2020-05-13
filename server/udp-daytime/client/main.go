package main

import (
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s host:port", os.Args[0])
	}

	address := os.Args[1]
	UDPAddress, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		log.Fatal(err)
	}

	// make a UDP connection request
	conn, err := net.DialUDP("udp", nil, UDPAddress)
	if err != nil {
		log.Fatal(err)
	}

	// send a message
	_, err = conn.Write([]byte("hello!"))
	if err != nil {
		log.Fatal(err)
	}

	var buf [512]byte // a buffer to hold the response
	n, err := conn.Read(buf[0:])
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(buf[0:n]))
}