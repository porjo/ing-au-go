// Package ingaugo provides a screenscraping interface to ING Australia Bank
package ingaugo

import (
	"os"

	"golang.org/x/exp/slog"
)

type Bank struct {
	wsURL  string
	logger *slog.Logger
}

type tokenResponse struct {
	Token string
}

// NewBank is used to initialize and return a Bank
// if websocketURL is not empty, the package will connect to browser instances listing at that location
// otherwise, the package will attempt to launch a local browser instance. It depends on 'google-chrome' executable being in $PATH
func NewBank(logger *slog.Logger, websocketURL string) (*Bank, error) {

	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout))
	}
	return &Bank{logger: logger, wsURL: websocketURL}, nil
}
