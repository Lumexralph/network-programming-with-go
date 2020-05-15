// Package main generates an RSA keys.
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"log"
	"os"
)

func main() {
	random := rand.Reader
	bits := 512
	key, err := rsa.GenerateKey(random, bits)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("private key primes %s\n", key.Primes[0].String())
	log.Printf("private key exponent %s\n", key.D.String())
	publicKey := key.PublicKey
	log.Printf("public key modulus %s\n", publicKey.N.String())
	log.Printf("public key exponent %s\n", publicKey.E)

	if err := saveGobKey("private.key", key); err != nil {
		log.Fatalln(err)
	}
	if saveGobKey("public.key", key.PublicKey); err != nil {
		log.Fatalln(err)
	}
	if savePEMKey("private.pem", key); err != nil {
		log.Fatalln(err)
	}
}

// the key is serialized/marshalled as gob
func saveGobKey(fileName string, key interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// marshaller/serializer
	e := gob.NewEncoder(file)
	return e.Encode(key)
}

// encode/serialize as a pem
func savePEMKey(fileName string, key *rsa.PrivateKey) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	return pem.Encode(file, privateKey)
}
