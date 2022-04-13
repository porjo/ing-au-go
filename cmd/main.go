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

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var bank ingaugo.Bank

func main() {

	accounts := make(arrayFlags, 0)

	wsURL := flag.String("ws-url", "", "WebSsocket URL e.g. ws://localhost:9222")
	clientNumber := flag.String("clientNumber", "", "Client number")
	accessPin := flag.String("accessPin", "", "Access pin")
	flag.Var(&accounts, "accountNumber", "Account number")
	days := flag.Int("days", 30, "Number of days of transactions")
	outputDir := flag.String("outputDir", "", "Directory to write CSV files. Defaults to current directory")
	flag.Parse()
	if *clientNumber == "" {
		fmt.Printf("-clientNumber is required\n\n")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *accessPin == "" {
		// check environmenta
		*accessPin = os.Getenv("ACCESS_PIN")
		if *accessPin == "" {
			fmt.Printf("-accessPin parameter or ACCESS_PIN environment variable is required\n\n")
			fmt.Println("Flags:")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
	if *outputDir != "" {
		info, err := os.Stat(*outputDir)
		if os.IsNotExist(err) {
			log.Fatalf("Directory %s does not exist", *outputDir)
		}
		if !info.IsDir() {
			log.Fatalf("%s is not a directory", *outputDir)
		}
	}

	// create a timeout as a safety net to prevent any infinite wait loops
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if *wsURL != "" {
		bank = ingaugo.NewBankWithWS(*wsURL)
	} else {
		bank = ingaugo.NewBank()
	}

	fmt.Printf("Fetching auth token...\n")
	token, err := bank.Login(ctx, *clientNumber, *accessPin)
	if err != nil {
		log.Fatal(err)
	}

	for _, acct := range accounts {
		err := GetTransactions(*days, acct, token, *outputDir)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetTransactions(days int, accountNumber, token, outputDir string) error {
	log.Printf("Fetching transactions for account %s\n", accountNumber)
	trans, err := bank.GetTransactionsDays(days, accountNumber, token)
	if err != nil {
		return err
	}

	file := accountNumber + ".csv"
	if outputDir != "" {
		file = outputDir + "/" + file
	}
	log.Printf("Writing CSV file %s\n", file)
	if err := os.WriteFile(file, trans, 0666); err != nil {
		return err
	}

	return nil
}
