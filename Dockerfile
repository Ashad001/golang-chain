# Use the official Go image to create a build artifact
FROM golang:1.22.3 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main main.go

# Debug step: list the contents of the /app directory
RUN ls -al /app

# Use a smaller base image to reduce the final image size
FROM alpine:latest

# Install ca-certificates to make HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Ensure the main executable has execute permissions
RUN chmod +x /root/main

# Debug step: list the contents of the /root directory again to confirm permissions
RUN ls -al /root

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ./root/main
