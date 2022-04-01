# ingaugo

[![Go Reference](https://pkg.go.dev/badge/github.com/porjo/ingaugo.svg)](https://pkg.go.dev/github.com/porjo/ingaugo)

A screenscraping interface to ING Australia Bank written in Go. It will drive a Chrome browser instance using the [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/).

## Usage

```Go
bank := ingaugo.NewBankWithWS(websocketURL)

ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
defer cancel()

token, err := bank.Login(ctx, *clientNumber, *accessPin)
if err != nil {
	log.Fatal(err)
}

log.Printf("token: %s\n", token)
```
The `websocketURL` refers to an already running instance of Chrome browser such as [headless-shell](https://hub.docker.com/r/chromedp/headless-shell/). Alternatively, use `NewBank()` and the package will attempt to launch Chrome browser locally by calling `google-chrome` executable.


## Credit

Based on https://github.com/adamroyle/ing-au-login
