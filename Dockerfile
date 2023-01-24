FROM golang:1.20rc2-alpine3.17

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY app.env .
COPY Makefile .
COPY .docker/certificates/ certificates
COPY src/ src/

RUN apk add --no-cache make

RUN go build -o server ./src/cmd/server.go

CMD ["./server"]
