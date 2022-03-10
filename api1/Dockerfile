ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY src .
COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR priceService
RUN go build

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY --from=builder /api/api1 .

EXPOSE 8080

ENTRYPOINT ["/api/api1"]
