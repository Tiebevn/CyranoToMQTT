# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy Go module files
COPY go.mod go.sum* ./
RUN go mod download

# Copy all Go source files
COPY *.go .

# Build the Go application
RUN go build -o cyrano-to-mqtt .

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/cyrano-to-mqtt .

# Expose the UDP port
EXPOSE 50103/udp

# Run the application
CMD ["./cyrano-to-mqtt"]
