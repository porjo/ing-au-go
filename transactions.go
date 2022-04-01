package ingaugo

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

const timeLayout = "2006-01-02T15:04:05-0700"
const exportTransactionsURL = "https://www.ing.com.au/api/ExportTransactions/Service/ExportTransactionsService.svc/json/ExportTransactions/ExportTransactions"

type transactionRequest struct {
	AuthToken       string `qs:"X-AuthToken"`
	AccountNumber   string
	Format          string
	FilterStartDate string
	FilterEndDate   string
	IsSpecific      bool
}

// GetTransactionsDays fetches transactions for the last x days. It takes an account number and auth token
// and returns CSV data
func (bank *Bank) GetTransactionsDays(days int, accountNumber, authToken string) (csv []byte, err error) {
	data := url.Values{}
	data.Set("X-AuthToken", authToken)
	data.Set("AccountNumber", accountNumber)
	data.Set("Format", "csv")
	data.Set("FilterStartDate", time.Now().AddDate(0, 0, -days).Format(timeLayout))
	data.Set("FilterEndDate", time.Now().AddDate(0, 0, 1).Format(timeLayout))
	data.Set("IsSpecific", "false")

	var c http.Client
	resp, err := c.PostForm(exportTransactionsURL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
