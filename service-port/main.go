/*
Services exists on host machines.

The IP Address is needed to locate the host machine but how will we
locate the services or distinguish the services on the host machines?
Ports are used to distinguish each service, i.e a service on an host machine
will use only one port or bind to a port and listen for connection or
request from that port
TCP, UDP and SCTP uses ports to send packets to another process or service
Port is an unsigned integer between 1 - 65535
Each service will allocate itself with one or more of these ports
Unix system, the commonly used ports are listed in the file /etc/services
Telnet - 23, DNS - 53, SSH - 22, HTTP 80, HTTPS - 443 etc
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s network-type service \n", os.Args[0])
	}

	network, service := os.Args[1], os.Args[2]
	port, err := net.LookupPort(network, service)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s is listening on port: %d\n", service, port)
}
