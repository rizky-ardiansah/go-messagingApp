# Build stage
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -s' -o simple-messaging-app

# Final stage - minimal image
FROM alpine:latest

# Install ca-certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/simple-messaging-app .
COPY --from=builder /app/.env .

# Expose ports
EXPOSE 4000
EXPOSE 8080

# Run the application
CMD ["./simple-messaging-app"]