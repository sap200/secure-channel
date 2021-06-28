// Package server contains the initial code
// to launch the server and handle the handshake packets
package server

import (
	"bufio"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/sap200/secure-channel/generator"
	"github.com/sap200/secure-channel/packets"
	"github.com/sap200/secure-channel/store"
	"github.com/sap200/secure-channel/utils"
)

// Starts listening on a port for incoming connections
func LaunchServer(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		// handle error
		panic(err)
	}
	fmt.Println("🎧 Listening for incoming connections on", ln.Addr().String())
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("Connection Terminated:", err)
		}

		go handleConnection(conn)
	}
}

// Handle connection handles the initial handshake packets sent for establishing secure channel
// once the handshake is compelted and secure connection is established
// handle connection launches 2 go routines one for reading from and another for writing to the connection respectively.
func handleConnection(conn net.Conn) {
	fmt.Println("🔗 Got an incoming connection from", conn.RemoteAddr().String())
	fmt.Println("🚧 Establishing secure connection with", conn.RemoteAddr().String())
	fmt.Println("🤝 Handshaking")

	// read the incoming public key value
	b, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		// unable to read public key
		fail(conn, "Failed: Unable to read from connection")
		return
	}
	// unmarshall the packet
	var synPack packets.SynPacket
	err = json.Unmarshal([]byte(b), &synPack)
	if err != nil {
		fail(conn, "Failed: Unable to marshal")
		return
	}

	// store it corresponding to its address
	store.Store[conn.RemoteAddr().String()] = synPack.PubKey

	// send him acknowledgement packet that contains your public key
	// send him your public key
	// he stores it in key value store
	ackPacket := packets.NewAckPacket(packets.AckSuccess, generator.PublicKey, "Success: Read my public key")
	bs, err := ackPacket.Marshall()
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(conn, bs)
	if err != nil {
		panic(err)
	}

	// Line open start communicating
	// line is on communicate using bufio
	fmt.Println("🔒 your channel is now secured using RSA cryptography")
	fmt.Println()
	go utils.Read(conn)
	go utils.Write(conn)

}

// When the handshake fails fail writes the acknowledgement fail packet to the connection.
func fail(conn net.Conn, message string) {
	var x rsa.PublicKey
	ackPacket := packets.NewAckPacket(packets.AckFail, x, message)
	bs, err := ackPacket.Marshall()
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, bs)
	conn.Close()
}
