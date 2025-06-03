FROM golang:1.23.3 AS builder

WORKDIR /build

ADD go.mod . 
ADD go.sum .

RUN go mod download

ADD . .

RUN go build -o mulus cmd/mulus/main.go

COPY build/mulus/mulus.yaml .


# FROM alpine:3.22
FROM ubuntu:24.04

WORKDIR /app
COPY --from=builder /build/mulus .
COPY --from=builder /build/mulus.yaml .

ENTRYPOINT [ "./mulus" ]
