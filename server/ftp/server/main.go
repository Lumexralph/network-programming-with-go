// Package main is an implentation of an FTP server.
package main

import (
	"log"
	"net"
	"os"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	address := ":5001"
	TCPAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.ListenTCP("tcp", TCPAddr)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
			return
		}

		s := string(buf[0:n])
		// decode the client request
		if s[0:2] == CD {
			chdir(conn, s[3:])
		} else if s[0:3] == DIR {
			listDir(conn)
		} else if s[0:3] == PWD {
			pwd(conn)
		}
	}
}

func chdir(conn net.Conn, s string) {
	if err := os.Chdir(s); err == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	conn.Write([]byte(s))
}

func listDir(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		conn.Write([]byte(err.Error()))
		return
	}

	for _, name := range names {
		conn.Write([]byte(name + "\r\n"))
	}
}
