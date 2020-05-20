// Package main is the implementation of simple proxy.
// Go uses the Transport field of the Request struct to
// to take a function that returns the URL of the proxy.
//
// Using Basic Authentication (username:password) to access the proxy.
// The authentication is encoded as Base64 before sent over the wire.
package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const basicAuth = "lumexralph:12345" // username:password

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s %s %s\n", os.Args[0], "http://proxy-host:port", "http://host:port/page")
	}

	proxyStr := os.Args[1]
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	destRawURL := os.Args[2]
	url, err := url.Parse(destRawURL)
	if err != nil {
		if err == io.EOF {
			return
		}

		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	// encode the authentication details
	auth := base64.StdEncoding.EncodeToString([]byte(basicAuth))
	// Transport for the proxy
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("GET", url.String(), nil)
	// proxy authentication
	req.Header.Add("Proxy-Authorization", auth)
	// the representation of the request over the wire
	dump, err := httputil.DumpRequest(req, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}

	fmt.Printf("dump: %s\n", string(dump))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(4)
	}
	defer resp.Body.Close()
	fmt.Println("Read ok")

	if resp.StatusCode != http.StatusAccepted || resp.Status != "200 OK" {
		fmt.Fprintln(os.Stderr, resp.Status)
		os.Exit(5)
	}

	fmt.Println("Response OK")

	var buf [512]byte
	for {
		// read from the response body to a buffer
		n, err := resp.Body.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(string(buf[0:n]))
	}
}
