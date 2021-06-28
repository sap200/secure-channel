package utils

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/sap200/secure-channel/generator"
	"github.com/sap200/secure-channel/packets"
	"github.com/sap200/secure-channel/store"
)

var Green = "\033[32m"
var Reset = "\033[0m"
var Blue = "\033[36m"

func Read(conn net.Conn) {
	for {
		b, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			// unable to read public key
			continue
		}

		var msg packets.MsgPacket
		// read a searilized string that is marshalled
		// unmarshall it into MsgPacket
		err = json.Unmarshal([]byte(b), &msg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// decrypt the message
		plain, err := rsa.DecryptPKCS1v15(rand.Reader, &generator.PrivateKey, msg.Cipher)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// verify the signature
		pubKey := store.Store[conn.RemoteAddr().String()]
		// hash the message
		hasher := sha512.New()
		hasher.Write(plain)
		hash := hasher.Sum(nil)
		err = rsa.VerifyPKCS1v15(&pubKey, crypto.SHA512, hash, msg.Signature)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Warning: Message is not verified..")
			continue
		}
		fmt.Println(Reset + Green + "" + conn.RemoteAddr().String() + ": " + string(plain) + Blue)

		// fmt.Print("\n" + conn.RemoteAddr().String() + " >>> " + b)
		// fmt.Print("you >>> ")
	}
}

func Write(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(Reset + Blue)
		text, _ := reader.ReadString('\n')
		fmt.Println(Reset + Green)

		// hash the text
		hasher := sha512.New()
		hasher.Write([]byte(text))
		hash := hasher.Sum(nil)
		// sign the text

		sign, err := rsa.SignPKCS1v15(rand.Reader, &generator.PrivateKey, crypto.SHA512, hash)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// encrypt the text
		pubKey := store.Store[conn.RemoteAddr().String()]
		cipher, err := rsa.EncryptPKCS1v15(rand.Reader, &pubKey, []byte(text))
		if err != nil {
			fmt.Println(err)
			continue
		}

		// form a message packet
		msg := packets.NewMsgPacket(cipher, sign)
		serialized, err := msg.Marshall()
		if err != nil {
			fmt.Println("Seems like channel is compromised")
			log.Fatalln(err)
		}

		// Marshall message packet and write
		io.WriteString(conn, serialized)
	}
}
