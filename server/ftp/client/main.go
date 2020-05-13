// Package main is the ftp client implementation.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	// strings used by the user interface.
	uiDIR  = "dir"
	uiCD   = "cd"
	uiPWD  = "pwd"
	uiQUIT = "quit"

	// strings used across the network
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	conn, err := net.Dial("tcp", ":5001")
	if err != nil {
		log.Fatalln(err)
	}

	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		// remove trailing space
		line = strings.TrimRight(line, "\t\r\n")
		if err != nil {
			log.Fatalln(err)
		}

		// split into command and argument
		strs := strings.SplitN(line, " ", 2)
		// decode the user request produced in stdin
		switch strs[0] {
		case uiDIR:
			dirRequest(conn)
		case uiCD:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}

			fmt.Println("CD\"", strs[1], "\"")
			cdRequest(conn, strs[1])
		case uiPWD:
			pwdRequest(conn)
		case uiQUIT:
			conn.Close()
		default:
			fmt.Println("unknown command")
		}
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + " "))
	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		// read till we hit a blank line
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Fatalln(err)
		}

		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512]byte
	n, err := conn.Read(response[0:])
	if err != nil {
		log.Fatalln(err)
	}

	if s := string(response[0:n]); s != "OK" {
		fmt.Println("failed to change dir")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512]byte
	n, err := conn.Read(response[0:])
	if err != nil {
		log.Fatalln(err)
	}

	s := string(response[0:n])
	fmt.Println("current directory \"" + s + "\"")
}
