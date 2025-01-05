FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o client main.go

# Use a minimal alpine image for the final container
FROM alpine:3.19

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/client .

CMD ["./client"]