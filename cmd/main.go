package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/porjo/ingaugo"
)

func main() {
	//	wsURL := flag.String("ws-url", "ws://localhost:9222", "WebSsocket URL")
	clientNumber := flag.String("clientNumber", "", "Client number")
	accessPin := flag.String("accessPin", "", "Access pin")
	accountNumber := flag.String("accountNumber", "", "Account number")
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

	// create a timeout as a safety net to prevent any infinite wait loops
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	//bank := ingaugo.NewBankWithWS(*wsURL)
	bank := ingaugo.NewBank()

	token, err := bank.Login(ctx, *clientNumber, *accessPin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("token: %s\n", token)

	if *accountNumber != "" {
		log.Printf("fetching transactions\n")
		trans, err := bank.GetTransactionsDays(30, *accountNumber, token)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("transactions\n%s\n", trans)
	}
}
