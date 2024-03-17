package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/porjo/ingaugo"
	"golang.org/x/exp/slog"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var bank *ingaugo.Bank
var logger *slog.Logger

func main() {

	accounts := make(arrayFlags, 0)

	wsURL := flag.String("ws-url", "", "WebSsocket URL e.g. ws://localhost:9222")
	clientNumber := flag.String("clientNumber", "", "Client number")
	accessPin := flag.String("accessPin", "", "Access pin")
	flag.Var(&accounts, "accountNumber", "Account number")
	days := flag.Int("days", 30, "Number of days of transactions")
	format := flag.String("format", "csv", "transaction output format (csv,ofx,qif)")
	outputDir := flag.String("outputDir", "", "Directory to write CSV files. Defaults to current directory")
	debug := flag.Bool("debug", false, "Output verbose logging")

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

	logOpts := slog.HandlerOptions{}
	if *debug {
		logOpts.Level = slog.LevelDebug
	} else {
		logOpts.Level = slog.LevelInfo
	}
	logger = slog.New(logOpts.NewTextHandler(os.Stdout))

	var err error
	bank, err = ingaugo.NewBank(logger, *wsURL)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Fetching auth token...")
	token, err := bank.Login(ctx, *clientNumber, *accessPin)
	if err != nil {
		log.Fatal(err)
	}

	if *debug {
		logger.Debug("token returned", "token", token)
	}

	for _, acct := range accounts {
		err := GetTransactions(*days, *format, acct, token, *outputDir)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetTransactions(days int, format string, accountNumber, token, outputDir string) error {
	logger.Info("Fetching transactions for account", "accountNumber", accountNumber)
	var f ingaugo.Format
	switch format {
	case "ofx":
		f = ingaugo.OFX
	case "qif":
		f = ingaugo.QIF
	default:
		f = ingaugo.CSV
	}
	trans, err := bank.GetTransactionsDays(days, f, accountNumber, token)
	if err != nil {
		return err
	}

	file := accountNumber + "." + format
	if outputDir != "" {
		file = outputDir + "/" + file
	}
	logger.Info("Writing transaction file", "file", file)
	if err := os.WriteFile(file, trans, 0666); err != nil {
		return err
	}

	return nil
}
