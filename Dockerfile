# Build stage
FROM docker.io/library/golang:alpine AS build-env

COPY . /go/ingaugo
WORKDIR /go/ingaugo/cmd

RUN apk update && \
    apk upgrade && \
	apk add git

RUN CGO_ENABLED=0 go build -o ingaugo

FROM docker.io/chromedp/headless-shell

RUN apt update && apt install ca-certificates -y

COPY --from=build-env /go/ingaugo/cmd /app/ingaugo

ENTRYPOINT ["/app/ingaugo/ingaugo"]
