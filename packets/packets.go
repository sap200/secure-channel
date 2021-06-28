package packets

import (
	"crypto/rsa"
	"encoding/json"
)

type Packet interface {
	Marshall() (string, error)
}

type AckPacket struct {
	AckStatus int64         `json:"ack_status"`
	PubKey    rsa.PublicKey `json:"pub_key"`
	Message   string        `json:"message"`
}

func NewAckPacket(ackStatus int64, pubKey rsa.PublicKey, message string) AckPacket {
	a := AckPacket{
		AckStatus: ackStatus,
		PubKey:    pubKey,
		Message:   message,
	}

	return a
}

func (a AckPacket) Marshall() (string, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}

type SynPacket struct {
	PubKey  rsa.PublicKey `json:"pub_key"`
	Message string        `json:"message"`
}

func NewSynPacket(pubKey rsa.PublicKey, message string) SynPacket {
	s := SynPacket{
		PubKey:  pubKey,
		Message: message,
	}

	return s
}

func (s SynPacket) Marshall() (string, error) {
	bs, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}

type MsgPacket struct {
	Cipher    []byte `json:"cipher"`
	Signature []byte `json:"signature"`
}

func (m MsgPacket) Marshall() (string, error) {
	bs, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(bs) + "\n", nil
}

func NewMsgPacket(cipher []byte, signature []byte) MsgPacket {
	m := MsgPacket{
		Cipher:    cipher,
		Signature: signature,
	}

	return m

}
