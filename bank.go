// Package ingaugo provides a screenscraping interface to ING Australia Bank
package ingaugo

type Bank struct {
	wsURL string
}

type tokenResponse struct {
	Token string
}

const loginURL string = "https://www.ing.com.au/securebanking/"

// NewBank is used to initialize and return a Bank that works by launching a
// a local browser instance. It depends on 'google-chrome' executable being in $PATH
func NewBank() Bank {
	return Bank{}
}

// NewBankWithWS initalises and returns a Bank that will attempt to
// connect to a browser via websocket URL
func NewBankWithWS(websocketURL string) Bank {
	return Bank{wsURL: websocketURL}
}
