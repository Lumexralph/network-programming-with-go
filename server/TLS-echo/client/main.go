// Package main contains the implementation of an echo server client,
// it makes a sure connection through TLS to the server.
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "usage: %s host:port\n", os.Args[0])
			os.Exit(1)
		}
	}
	// TODO: solve the client TLS issue
	// load the public and private key
	cert, err := tls.LoadX509KeyPair("../../../encryption/x-509/gen-cert/lumexralph.github.io.pem", "../../../encryption/x-509/gen-cert/private.pem")
	addr := os.Args[1]
	// create TLS configuration.
	config := tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	// make a TLS connection
	conn, err := tls.Dial("tcp", addr, &config)
	if err != nil {
		log.Fatalln(err)
	}

	for n := 0; n < 15; n++ {
		log.Println("Writing...")
		// send a message to the connection stream, to the server.
		conn.Write([]byte("Hello " + string(n+48)))

		var buf [512]byte
		// read the response and store in memory using the buffer.
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(buf[0:n]))
	}
}
