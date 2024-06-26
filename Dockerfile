FROM golang:1.17 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
FROM scratch
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
