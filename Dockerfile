FROM golang:1.21 as builder

# Install Go if it's not already installed
RUN if ! [ -x "$(command -v go)" ]; then echo "Installing Go..." && \
    apk add --no-cache curl && \
    curl -L "https://golang.org/dl/go1.21.linux-amd64.tar.gz" | tar -xz && \
    mv go /usr/local && \
    rm -rf /var/cache/apk/*; fi

WORKDIR /app

COPY . /app

RUN go build -o main .
RUN chmod +x /app/main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/
RUN ls -l /app
RUN pwd
CMD ["./main"]