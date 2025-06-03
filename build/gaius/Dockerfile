FROM golang:1.23.3 AS builder

WORKDIR /build

ADD go.mod . 
ADD go.sum .

RUN go mod download

ADD . .

RUN go build -o gaius cmd/gaius/main.go

COPY build/gaius/gaius.yaml .


# FROM alpine:3.22
FROM ubuntu:24.04

WORKDIR /app
COPY --from=builder /build/gaius .
COPY --from=builder /build/gaius.yaml .

ENTRYPOINT [ "./gaius" ]
