FROM golang:1.22.0 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build ./cmd/main.go

FROM alpine:3.19.1 AS production
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT ["./main"]