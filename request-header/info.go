package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s hostname:port\n", os.Args[0])
	}

	serviceAddr := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serviceAddr)
	if err != nil {
		log.Fatal(err)
	}

	// get a connection with the server
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	// make a HEAD request
	if _, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")); err != nil {
		log.Fatalln(err)
	}

	// read the response message from the connection stream
	response, err := ioutil.ReadAll(conn)
	fmt.Printf("response message: %s\n", response)
}
