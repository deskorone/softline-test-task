FROM golang:1.16-alpine

WORKDIR /app

COPY internal ./internal/

COPY cmd ./cmd/

COPY ./go.mod .

COPY ./config.yaml .

COPY pkg ./pkg

EXPOSE 8080

RUN go mod tidy

RUN go build -o server ./cmd/main.go

CMD ["./server"]

