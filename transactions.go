package ingaugo

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	timeLayout            = "2006-01-02T15:04:05-0700"
	exportTransactionsURL = "https://www.ing.com.au/api/ExportTransactions/Service/ExportTransactionsService.svc/json/ExportTransactions/ExportTransactions"

	// Make Go HTTP client user-agent match headless-shell user-agent
	userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36"

	CSV Format = "csv"
	OFX Format = "ofx"
	QIF Format = "qif"
)

type Format string

/*
type transactionRequest struct {
	AuthToken       string `qs:"X-AuthToken"`
	AccountNumber   string
	Format          string
	FilterStartDate string
	FilterEndDate   string
	IsSpecific      bool
}
*/

// GetTransactionsDays fetches transactions for the last x days. It takes an account number and auth token
func (bank *Bank) GetTransactionsDays(days int, format Format, accountNumber, authToken string) ([]byte, error) {
	data := url.Values{}
	data.Set("X-AuthToken", authToken)
	data.Set("AccountNumber", accountNumber)
	data.Set("Format", string(format))
	data.Set("FilterStartDate", time.Now().AddDate(0, 0, -days).Format(timeLayout))
	data.Set("FilterEndDate", time.Now().AddDate(0, 0, 1).Format(timeLayout))
	data.Set("IsSpecific", "false")

	c := &http.Client{}

	bank.logger.Info("Fetching page", "url", exportTransactionsURL)
	req, err := http.NewRequest("POST", exportTransactionsURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		bank.logger.Info("Response body", "body", string(body))
		return nil, fmt.Errorf("error fetching transactions. Status code: %d", resp.StatusCode)
	}

	return body, nil
}
