FROM golang:1.26 as builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLE=0 go build -o server

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080

CMD ["./server"]
