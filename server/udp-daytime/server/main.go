package main

import (
	"net"
	"time"

	"log"
)

func main() {
	addr := ":8000"
	UDPAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	// listen for connection
	conn, err := net.ListenUDP("udp", UDPAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		handleDayTime(conn)
	}
}

func handleDayTime(conn *net.UDPConn) {
	var buf [512]byte

	_, UDPAddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		log.Println(err)
		return
	}

	daytime := time.Now().String()
	_, err = conn.WriteToUDP([]byte(daytime), UDPAddr)
	if err != nil {
		log.Println(err)
		return
	}
}
