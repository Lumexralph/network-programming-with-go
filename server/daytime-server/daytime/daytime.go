package daytime

import (
	"log"
	"net"
	"time"
)

// CurrentTime will give the curent
func CurrentTime(conn net.Conn) {
	daytime := time.Now().String()
	conn.Write([]byte(daytime + "\n"))
	log.Println("responded to new client")
	// we are done with the client, time to close up.
	// and wait for the next call.
	if err := conn.Close(); err != nil {
		log.Println(err)
	}
}
