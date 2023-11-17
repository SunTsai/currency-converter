# Build stage
FROM golang:1.20.3-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY exchange_rates.json .
ENV GIN_MODE=release

EXPOSE 8080
CMD [ "/app/main" ]