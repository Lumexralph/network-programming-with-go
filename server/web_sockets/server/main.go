// Package main is echo server that uses web sockets
// to push message from the server to the client.
package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	fmt.Println("Echoing...")

	for n := 0; n < 10; n++ {
		msg := fmt.Sprintf("Hello %d", n)
		fmt.Printf("Sending message to client: %s\n", msg)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("can't send")
			break
		}

		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("can't receive")
			break
		}

		fmt.Printf("Message received back from client: %s\n", reply)
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	// you can make it a secure web socket by using TLS and X.509 Certificates.
	// client will connect using wss://localhost:5001 if not secured then
	// it will ws://localhost:5001
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
