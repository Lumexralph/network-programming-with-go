/*Package main will prepare an IP connection,
send a ping request to a host and get a reply.
You may need to have root access in order to run it successfully.
*/
package main

import (
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s host", os.Args[0])
	}

	addr := os.Args[1]
	IPAddr, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		log.Fatal(err)
	}

	// make an IP connection request using ICMP protocol
	conn, err := net.DialIP("ip4:icmp", nil, IPAddr)
	if err != nil {
		log.Fatal(err)
	}

	// buffer to hold messages/packets
	var msg [512]byte
	msg[0] = 8  // echo message
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum in entire message
	msg[3] = 0  // checksum in entire message
	msg[4] = 0  // arbitrary identifier 1
	msg[5] = 13 // arbitrary identifier 2
	msg[6] = 0  // arbitrary sequence number 1
	msg[7] = 30 // arbitrary sequence number 2
	len := 8

	// create checksum
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Read(msg[0:])
	if err != nil {
		log.Fatal(err)
	}

	log.Println("got a response")
	// check the packet
	if msg[5] == 13 {
		log.Println("packet identifiers matches")
	}
	if msg[7] == 30 {
		log.Println("sequence matches")
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}
