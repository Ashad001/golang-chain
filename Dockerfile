FROM golang:1.22.3 as builder

WORKDIR /app

COPY . .

# Build the Go application
RUN go build -o main ./main.go

# Use a minimal base image
FROM alpine:latest

WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/main /app/

# Ensure the binary has execution permissions
RUN chmod +x /app/main

# Verify the binary is in place and executable
RUN ls -l /app
RUN pwd

# Set the entrypoint to run the Go binary
CMD ["./main"]
