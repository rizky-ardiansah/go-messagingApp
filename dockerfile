FROM golang:1.24.0-alpine

WORKDIR /app

# Copy go.mod and go.sum for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o simple-messaging-app

# Set executable permissions
RUN chmod +x simple-messaging-app

# Expose ports
EXPOSE 4000
EXPOSE 8080

# Run the application
CMD ["./simple-messaging-app"]