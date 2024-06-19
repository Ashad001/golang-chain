# Build stage
FROM golang:1.22.3 as builder

WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Ensure the binary has execute permissions
RUN chmod +x main

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary has execute permissions (redundant but safe)
RUN chmod +x main

# Run the binary
CMD ["./main"]
