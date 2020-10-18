FROM golang:1.15.2-alpine3.12 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o gonca ./cmd/gonca/main.go

FROM scratch

COPY --from=builder /app/gonca .
COPY --from=builder /app/.config.yml .

CMD ["./gonca"]
