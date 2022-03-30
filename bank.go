package ingaugo

import (
	"context"

	dp "github.com/chromedp/chromedp"
)

type Bank struct {
	context context.Context
	cancel  context.CancelFunc

	clientNumber string
	accessPin    string
}

type tokenResponse struct {
	Token string
}

const loginURL string = "https://www.ing.com.au/securebanking/"

func NewBank(url string) Bank {
	b := Bank{}
	if url != "" {
		b.context, b.cancel = dp.NewRemoteAllocator(context.Background(), url)

	}
	if b.context == nil {
		b.context = context.Background()
	}
	b.context, b.cancel = dp.NewContext(b.context)
	return b
}
