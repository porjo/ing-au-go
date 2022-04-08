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
	days := flag.Int("days", 30, "Number of days of transactions")
	flag.Parse()
	if *clientNumber == "" {
		fmt.Printf("-clientNumber is required\n\n")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *accessPin == "" {
		// check environmenta
		*accessPin = os.Getenv("ING_ACCESS_PIN")
		if *accessPin == "" {
			fmt.Printf("-accessPin parameter or ING_ACCESS_PIN environment variable is required\n\n")
			fmt.Println("Flags:")
			flag.PrintDefaults()
			os.Exit(1)
		}
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

	if *accountNumber != "" {
		log.Printf("fetching transactions\n")
		trans, err := bank.GetTransactionsDays(*days, *accountNumber, token)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s\n", trans)
	}
}
