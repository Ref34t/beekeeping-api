# Use official Golang image as the base
FROM golang:1.17-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o beekeeping-api main.go

# Use a smaller base image for the final container
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/beekeeping-api .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./beekeeping-api"]