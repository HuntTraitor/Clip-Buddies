FROM golang:alpine

LABEL maintainer="Hunter Tratar"

RUN apk add --no-cache make gcc libc-dev git

WORKDIR /backend

RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

ENV PATH=$PATH:/go/bin

COPY go.mod go.sum ./
RUN go mod download