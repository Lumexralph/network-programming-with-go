// package main is an implementation of a daytime server
// according to https://tools.ietf.org/html/rfc868
//
// Error handling in the server as compared to a client is different.
// The server should run forever, so that if any error occurs with a client
// the server just ignores that client and carries on.

Excerpt From: Unknown. “Table of Contents”. Apple Books. 
package main

import (
	"log"
	"net"
	"time"
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
		// the server can only respond to one client at a time
		log.Println("waiting for new client connection")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		// respond to the client, we need to send bytes.
		time.Sleep(15 * time.Second) // simulate a delay
		conn.Write([]byte(daytime + "\n"))
		log.Println("responded to new client")
		// we are done with the client, time to close up.
		// and wait for the next call.
		if err := conn.Close(); err != nil {
			log.Println(err)
		}
	}
}
