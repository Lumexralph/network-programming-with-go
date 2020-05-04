// Package main has implementation of how to handle
// masking operations and find the network for that IP address.

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

type IPMask []byte

// IPV4Mask creates a mask (which can help to know the
// external and internal network address of a machine)
// from a 32 bits or 4 bytes IPV4 address
func IPV4Mask(a, b, c, d byte) IPMask {
	return IPMask{}
}

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

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Printf(`address is %s default mask length is %d
	leading ones count is %d mask is %x
	network is %s
	`, addr.String(), bits, ones, mask.String(), network.String())
}
