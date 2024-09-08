package main

import (
	"log"

	"github.com/Hamzenium/Distributed-CAS-Storage/p2p"
)

// t initializes a new TCP transport listening on port 3000 with p2p.NewTCPTransport(":3000").
// The ListenAndAccept method is called on the transport object, which presumably starts listening for and accepting incoming connections. If an error occurs, it logs the error and terminates the program.
// The select {} statement is used to block the main goroutine indefinitely, keeping the program running//
func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal((err))
	}
	select {}
}
