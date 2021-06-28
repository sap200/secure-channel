// package packets
// consists of the SYN packet, ACK packet and MSG Packet
// This packets helps in easy transnportation and decoding to establish a secure channel
package packets

import (
	"crypto/rsa"
	"encoding/json"
)

// packet is an interface
// any type implementing method Marshall belongs to type Packet
type Packet interface {
	Marshall() (string, error)
}

// This is acknowledgement packet
// It consists of 3 fields
// Acknowledgement status i.e. either fail or success
// Public key of the sender
// Message sent by the sender to further say what was error in acknowledgement
type AckPacket struct {
	AckStatus int64         `json:"ack_status"`
	PubKey    rsa.PublicKey `json:"pub_key"`
	Message   string        `json:"message"`
}

// Makes the New acknowledgement packet and returns it
func NewAckPacket(ackStatus int64, pubKey rsa.PublicKey, message string) AckPacket {
	a := AckPacket{
		AckStatus: ackStatus,
		PubKey:    pubKey,
		Message:   message,
	}

	return a
}

// Marshall marshalls the acknowledgement packet using json package of golang
func (a AckPacket) Marshall() (string, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}

// Syn Packet is used to send sender's public key with extraa greeting message
// This packet is usually sent by the client
type SynPacket struct {
	PubKey  rsa.PublicKey `json:"pub_key"`
	Message string        `json:"message"`
}

// NewSynPacket makes a new syn packet with given arguments
func NewSynPacket(pubKey rsa.PublicKey, message string) SynPacket {
	s := SynPacket{
		PubKey:  pubKey,
		Message: message,
	}

	return s
}

// Marshall marshalls using json package of golang
func (s SynPacket) Marshall() (string, error) {
	bs, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}

// type MsgPacket contains the cipher and signature
// signature needs to be validated and cipher needs to be decrypted
type MsgPacket struct {
	Cipher    []byte `json:"cipher"`
	Signature []byte `json:"signature"`
}

// New MsgPacket makes new msg packet
func NewMsgPacket(cipher []byte, signature []byte) MsgPacket {
	m := MsgPacket{
		Cipher:    cipher,
		Signature: signature,
	}

	return m
}

// Marshall marshalls the msg packet using existing json package of golang
func (m MsgPacket) Marshall() (string, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}
