FROM golang:1.20rc2-alpine3.17 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY app.env .
COPY Makefile .
COPY src/ src/

RUN apk add --no-cache make

RUN go build -o server ./src/cmd/server.go

############################################
FROM alpine:3.14

WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080

CMD ["./server"]