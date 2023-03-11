FROM golang:1.19-alpine

WORKDIR /app

COPY internal ./internal/

COPY cmd ./cmd/

COPY ./go.mod .

COPY ./config.yaml .

COPY pkg ./pkg

RUN go mod tidy

RUN go build -o server ./cmd/main.go

CMD ["./server"]

