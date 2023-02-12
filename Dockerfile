FROM golang:1.20rc2-alpine3.17 AS builder

WORKDIR /goapp

COPY go.mod .
COPY go.sum .
COPY app.env .
COPY app/ ./app

RUN go build -o server ./app/cmd/server.go

############################################
FROM alpine:3.14

WORKDIR /app
COPY --from=builder /goapp/server .
EXPOSE 8080

CMD ["./server"]