// package client dials the server and
// sends a syn packet and decodes the acknowledgement packet
// to establish a secure channel among them
package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/sap200/secure-channel/generator"
	"github.com/sap200/secure-channel/packets"
	"github.com/sap200/secure-channel/store"
	"github.com/sap200/secure-channel/utils"
)

var wg sync.WaitGroup

// Dial server dials a server with given address
// Then sends a Syn Packet to them
// Receives a Acknowledgement packet
// Decodes the ack packet
// If Ack Status is success
// It launches 2 go routines one for reading from the server
// and another for writing to the server.
func DialServer(address string) {
	fmt.Println("🚧 Establishing secure connection with", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		// handle error
		panic(err)
	}

	fmt.Println("🤝 Handshaking")

	// Form a syn packet
	s := packets.NewSynPacket(generator.PublicKey, "Check: Here is my public key check it out")
	bs, err := s.Marshall()
	if err != nil {
		panic(err)
	}
	// write bs to the connection
	_, err = io.WriteString(conn, bs)
	if err != nil {
		panic(err)
	}
	// here for the ack packet
	// read from conn and decode it to ack packet
	// read the incoming public key value
	b, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		// unable to read public key
		fmt.Println("Unable to read from connection")
		conn.Close()
	}

	// unmarshall the packet
	var ackPack packets.AckPacket
	err = json.Unmarshal([]byte(b), &ackPack)
	if err != nil {
		// unexpected close connection
		fmt.Println("Unable to unmarshal ack packet")
		conn.Close()
	}

	// check if ack packet is alright
	if ackPack.AckStatus == packets.AckFail {
		// unable to connect
		fmt.Println("Ack Failed..")
		conn.Close()
	}

	// store the public key
	store.Store[conn.RemoteAddr().String()] = ackPack.PubKey

	// start communicating
	fmt.Println("🔒  your channel is now secured using RSA cryptography")
	fmt.Println()
	wg.Add(2)

	go utils.Read(conn)
	go utils.Write(conn)

	wg.Wait()
}
