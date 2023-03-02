FROM golang:1.20rc2-alpine3.17 AS builder

WORKDIR /goapp

COPY go.mod .
COPY go.sum .
COPY app.env .
COPY app/ ./app

RUN go build -o main ./app/cmd/main.go

############################################
FROM alpine:3.14

WORKDIR /app
COPY --from=builder /goapp/main .
EXPOSE 8080

CMD ["./main"]