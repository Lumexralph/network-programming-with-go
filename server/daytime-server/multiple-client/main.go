// package main is an implementation of a daytime server
// according to https://tools.ietf.org/html/rfc868
//
// Error handling in the server as compared to a client is different.
// The server should run forever, so that if any error occurs with a client
// the server just ignores that client and carries on.
package main

import (
	"log"
	"net"

	"go-network-programming/daytime-server/daytime"
)

func main() {
	address := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		log.Fatal(err)
	}

	// bind the server to a port on the host machine
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("could not bind to port: %s\n", err)
	}

	// server has to be in a continuous state
	// listening for connection requests.
	for {
		// on an accept request or call from the client
		// create a connection object that is returned.
		// it will block on accept() till a connection request is sent.
		log.Println("waiting for new client connection")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// handle more client requests in new goroutines
		go daytime.CurrentTime(conn)
	}
}
