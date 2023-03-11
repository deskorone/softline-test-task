FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o server ./cmd/main.go

CMD ["./server"]

