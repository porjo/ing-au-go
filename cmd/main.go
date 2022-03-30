package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/porjo/ingaugo"
)

func main() {
	wsURL := flag.String("ws-url", "ws://localhost:9222", "WebSsocket URL")
	clientNumber := flag.String("clientNumber", "", "Client number")
	accessPin := flag.String("accessPin", "", "Access pin")
	flag.Parse()
	if *clientNumber == "" {
		fmt.Println("clientNumber is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *accessPin == "" {
		fmt.Println("accessPin is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	bank := ingaugo.NewBank(*wsURL)

	token, err := bank.Login(*clientNumber, *accessPin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("token: %s\n", token)
}
