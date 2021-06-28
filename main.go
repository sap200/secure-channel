// package main parses the command line args
// and routes to appropriate modules accordingly
// It acts like a handler
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/sap200/secure-channel/client"
	"github.com/sap200/secure-channel/generator"
	"github.com/sap200/secure-channel/server"
	"github.com/sap200/secure-channel/store"
	"github.com/sap200/secure-channel/utils"
)

// Initializes the private key and in memory store
func init() {
	generator.PrivateKey = generator.NewGenerator()
	generator.PublicKey = generator.PrivateKey.PublicKey
	// initialize store
	store.Store = store.InitStore()

	// dont change colors in windows it doesn't supports
	if runtime.GOOS == "windows" {
		utils.Green = ""
		utils.Reset = ""
		utils.Blue = ""
	}
}

// parses the flag
func main() {

	command := flag.String("command", "", "command: either server or client")
	ip := flag.String("ip", "", "ip: if command is server then launch ip, if command is client then server's ip  to connect")
	port := flag.Uint("port", 0, "port: if command is server then launch port, if command is client then server's port to connect")

	flag.Parse()

	if *command != "server" && *command != "client" {
		log.Fatalln("command should either be 'server' or 'client'")
	}

	if *ip == "" {
		log.Fatalln("Please provide a valid ip")
	}

	if *port == 0 {
		log.Fatalln("Please provide a valid port, usually > 2000")
	}

	// form the address
	address := fmt.Sprintf("%s:%v", *ip, *port)

	switch *command {
	case "server":
		// launch server
		server.LaunchServer(address)
	case "client":
		client.DialServer(address)
	default:
		log.Fatalln("Invalid command")
	}

	// This one cleans up the terminal color in linux on press of ctrl+c
	// not tested
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		<-c
		fmt.Print(utils.Reset)
	}()

}
