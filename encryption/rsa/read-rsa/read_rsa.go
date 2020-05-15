// Package main will load and decrypt an RSA generated key.
package main

import (
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	var key rsa.PrivateKey
	if err := loadKey("../gen-rsa/private.key", &key); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("private key primes: %s \n --- %s\n",
		key.Primes[0].String(), key.Primes[0].String())
	fmt.Printf("private key exponent %s\n", key.D.String())

	var publicKey rsa.PublicKey
	if err := loadKey("../gen-rsa/public.key", &publicKey); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("public key modulus %s\n", publicKey.N.String())
	fmt.Printf("public key exponent %s\n", publicKey.E)
}

func loadKey(fileName string, key interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// deserialize/unmarshal the gob
	d := gob.NewDecoder(file)
	return d.Decode(key)
}
