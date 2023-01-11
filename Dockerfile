FROM golang:1.20rc2-alpine3.17

WORKDIR /app
COPY . .

RUN go build -o main .

CMD ["./main"]
