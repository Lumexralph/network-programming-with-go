// Package main is showing how to parse an IPv4 or IPv6 address
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// ensure a command line argument is provided
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	arg := os.Args[1]
	addr := net.ParseIP(arg)
	if addr == nil {
		log.Fatal("invalid IP address")
	}

	fmt.Printf("the address is %s\n", addr.String())
}
