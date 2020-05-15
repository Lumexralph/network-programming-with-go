// Package main will decrypt an x509 certificate.
package main

import (
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

func main() {
	certCerFile, err := os.Open("../gen-cert/lumexralph.github.io.cer")
	if err != nil {
		log.Fatalln(err)
	}
	// ensure it is bigger than the file to be
	// able to read the bytes in the cer file.
	certBytes := make([]byte, 1000)
	count, err := certCerFile.Read(certBytes)
	if err != nil {
		log.Fatalln(err)
	}

	// trim the file to the actual number of bytes
	// that was read from certCerFile.
	cert, err := x509.ParseCertificate(certBytes[0:count])
	if err != nil {
		log.Fatalln(err)
	}

	// the details of the public certificate
	fmt.Printf("serial number: %d\n", cert.SerialNumber)
	fmt.Printf("country: %s\n", cert.Subject.Country[0])
	fmt.Printf("organization: %s\n", cert.Subject.Organization[0])
	fmt.Printf("common name: %s\n", cert.Subject.CommonName)
	fmt.Printf("DNS Names: %s\n", cert.DNSNames)
}
