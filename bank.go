// Package ingaugo provides a screenscraping interface to ING Australia Bank
package ingaugo

import "log"

type Bank struct {
	wsURL string
}

type customLog struct {
	debugLog bool
}

var clog customLog

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

// SetDebug turns on/off verbose logging to stderr
func SetDebug(state bool) {
	clog.debugLog = state
}

func (l customLog) Printf(format string, v ...interface{}) {
	if l.debugLog {
		log.Printf(format, v)
	}
}
func (l customLog) Println(msg string) {
	l.Printf("%v\n", msg)
}
