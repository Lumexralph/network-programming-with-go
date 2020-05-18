// Package main is the implementation of a TLS
// secured echo server.
package main

import (
	"crypto/rand"
	"crypto/tls"
	"log"
	"net"
	"time"
)

func main() {
	// load the public and private key
	cert, err := tls.LoadX509KeyPair("../../../encryption/x-509/gen-cert/lumexralph.github.io.pem", "../../../encryption/x-509/gen-cert/private.pem")
	if err != nil {
		log.Fatalln(err)
	}

	// create TLS configuration.
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = func() time.Time {
		return time.Now()
	}
	config.Rand = rand.Reader

	address := ":5000"
	list, err := tls.Listen("tcp", address, &config)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listening...")

	for {
		conn, err := list.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Accepted!")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		log.Println("Trying to read..")
		// read from the connection stream into the buffer
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
		}

		// write back to the stream what was sent from the buffer.
		_, err = conn.Write(buf[0:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}
