package ingaugo

import (
	"io"
	"net/http"
	"time"

	"github.com/ajg/form"
)

const timeLayout = "2006-01-02T15:04:05Z"
const exportTransactionsURL = "https://www.ing.com.au/api/ExportTransactions/Service/ExportTransactionsService.svc/json/ExportTransactions/ExportTransactions"

type TransactionRequest struct {
	AuthToken       string `form: "X-AuthToken"`
	AccountNumber   string
	Format          string
	FilterStartDate string
	FilterEndDate   string
	IsSpecific      bool
}

func (bank *Bank) FetchLast30Days(accountNumber, authToken string) (csv []byte, err error) {
	data := TransactionRequest{
		AuthToken:       authToken,
		AccountNumber:   accountNumber,
		Format:          "csv",
		FilterStartDate: time.Now().AddDate(0, 0, -30).Format(timeLayout),
		FilterEndDate:   time.Now().AddDate(0, 0, 1).Format(timeLayout),
		IsSpecific:      false,
	}

	var c http.Client
	vals, err := form.EncodeToValues(data)
	if err != nil {
		return nil, err
	}
	resp, err := c.PostForm(exportTransactionsURL, vals)
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
