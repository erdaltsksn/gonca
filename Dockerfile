FROM golang:1.15.2-alpine3.12 AS builder

WORKDIR /app

COPY . .

RUN go build -o gonca ./cmd/gonca/main.go

FROM scratch

COPY --from=builder /app/gonca .

CMD ["./gonca"]
