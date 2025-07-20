FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./target/simple-blockchain ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/target/simple-blockchain ./simple-blockchain

ENV MODE=prod
ENV PORT=8000

EXPOSE $PORT

CMD ["/app/simple-blockchain"]

