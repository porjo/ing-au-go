# ingaugo

[![Go Reference](https://pkg.go.dev/badge/github.com/porjo/ingaugo.svg)](https://pkg.go.dev/github.com/porjo/ingaugo)

A screenscraping interface to ING Australia Bank written in Go. It will drive a Chrome browser instance using the [Chrome DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/).

## Usage

```Go
bank, err = ingaugo.NewBank(logger, *wsURL)
if err != nil {
	log.Fatal(err)
}

ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
defer cancel()

token, err := bank.Login(ctx, *clientNumber, *accessPin)
if err != nil {
	log.Fatal(err)
}

log.Printf("token: %s\n", token)
```
`wsURL` refers to an already running instance of Chrome browser such as [headless-shell](https://hub.docker.com/r/chromedp/headless-shell/). If `wsURL` is nil then the package will attempt to launch Chrome browser locally by calling `google-chrome` executable.

### Docker image

A docker image is available which provides a cli:

```
$ docker pull ghcr.io/porjo/ingaugo:latest

$ docker run --rm ingaugo
-clientNumber is required

Flags:
  -accessPin string
    	Access pin
  -accountNumber value
    	Account number
  -clientNumber string
    	Client number
  -days int
    	Number of days of transactions (default 30)
  -debug
    	Output verbose logging
  -outputDir string
    	Directory to write CSV files. Defaults to current directory
  -ws-url string
    	WebSsocket URL e.g. ws://localhost:9222
``` 


## Credit

Based on https://github.com/adamroyle/ing-au-login
