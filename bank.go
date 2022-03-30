// Package ingaugo provides a screenscraping interface to ING Australia Bank
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

// NewBank takes an optional url which refers to a browser's websocket address.
// If no url is supplied, it will attempt to launch a local browser instance.
func NewBank(websocketUrl string) Bank {
	b := Bank{}
	if websocketUrl != "" {
		b.context, b.cancel = dp.NewRemoteAllocator(context.Background(), websocketUrl)

	}
	if b.context == nil {
		b.context = context.Background()
	}
	b.context, b.cancel = dp.NewContext(b.context)
	return b
}
