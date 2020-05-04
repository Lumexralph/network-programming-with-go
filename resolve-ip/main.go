// Package main illustrates a way to perform a DNS lookup
// for any URL. Also looks up the possible addresses for the URL.
//
// Example: resolve-ip www.github.com
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s hostname\n", os.Args[0])
	}

	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		log.Fatalf("resolution error %s", err)
	}

	fmt.Printf("resolution address %s\n", addr)

	addrs, err := net.LookupHost(name)
	if err != nil {
		log.Fatalf("lookup address error %s", err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
