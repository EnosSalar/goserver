# Start from the official Go image
FROM golang:1.25-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app (statically linked binary)
RUN go build -o myapp

# Use a minimal image for running the binary
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the port your app uses
EXPOSE 8080

# Command to run the binary
CMD ["./myapp"]
