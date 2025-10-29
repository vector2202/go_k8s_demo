FROM golang:1.25.1 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download || true

COPY . .
RUN go build -o server .

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 9090
CMD ["./server"]
