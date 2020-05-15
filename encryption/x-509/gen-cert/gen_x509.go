// Package main generates a self-signed X.509 certificate.
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader
	var key rsa.PrivateKey
	if err := loadKey("../../rsa/gen-rsa/private.key", &key); err != nil {
		log.Fatalln(err)
	}

	currentTime := time.Now()
	nextTime := currentTime.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000) // next one year
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:      []string{"Nigeria"},
			Organization: []string{"Lumexralph Tech"},
			CommonName:   "lumexralph",
		},
		NotBefore:             currentTime,
		NotAfter:              nextTime,
		SubjectKeyId:          []byte{1, 2, 3, 4},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"lumexralph.github.io", "localhost"},
	}

	cert, err := x509.CreateCertificate(random, template, template, &key.PublicKey, &key)
	if err != nil {
		log.Fatalln(err)
	}

	certCerFile, err := os.Create("lumexralph.github.io.cer")
	if err != nil {
		log.Fatalln(err)
	}
	defer certCerFile.Close()

	_, err = certCerFile.Write(cert)
	if err != nil {
		log.Fatalln(err)
	}

	certPEMFile, err := os.Create("lumexralph.github.io.pem")
	if err != nil {
		log.Fatalln(err)
	}
	defer certPEMFile.Close()

	pem.Encode(certPEMFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})

	keyPEMFile, err := os.Create("private.pem")
	if err != nil {
		log.Fatalln(err)
	}
	defer keyPEMFile.Close()

	pem.Encode(keyPEMFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key),
	})
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
