FROM golang:1.20rc2-alpine3.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY app.env .
COPY Makefile .
COPY src/ src/

RUN go build -o server ./src/cmd/server.go

CMD ["./server"]
