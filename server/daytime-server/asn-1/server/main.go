/* Package main is the implementation DaytimeServer
using ASN1 encoded for communication.
*/
package main

import (
	"encoding/asn1"
	"log"
	"net"
	"time"
)

func main() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now()
		// responded with ASN encoded data
		data, err := asn1.Marshal(daytime)
		if err != nil {
			log.Fatal(err)
		}

		conn.Write(data)
		conn.Close()
	}
}
