# Start from the official Golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app with statically linked libraries
RUN CGO_ENABLED=0 go build -ldflags "-s -w -extldflags '-static'" -o main .

# Start a new stage from Debian base image
FROM debian:buster-slim

# Create a new user to run the application
RUN useradd -u 10001 appuser

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Change the ownership of the binary to the non-root user
RUN chown appuser:appuser ./main

# Expose port 8080 to the outside world
EXPOSE 8080

# Set the user to run the application
USER appuser

# Command to run the executable
CMD ["/app/main"]
