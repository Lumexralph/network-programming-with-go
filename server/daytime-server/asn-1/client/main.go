// Package main is the implementation of the ASN1 daytime server client.
// reading asn1 encoded data.
package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	var newtime time.Time
	_, err = asn1.Unmarshal(result, &newtime)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After marshal/unmarshal: ", newtime.String())
}
