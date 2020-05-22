// Package main is the websocket client for the echo server.
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("usage: ", os.Args[0], "ws://host:port")
	}

	url := os.Args[1]
	conn, err := websocket.Dial(url, "", "http://localhost")
	if err != nil {
		log.Fatalln(err)
	}

	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				break
			}

			fmt.Printf("couldn't receive message: %v", err)
			break
		}
		fmt.Printf("received message from server %s\n", msg)

		// return message
		err = websocket.Message.Send(conn, msg+" Lumex")
		if err != nil {
			fmt.Println("could not return message")
			break
		}
	}
}
