FROM golang:1.22

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /golangchain

CMD ["/golangchain"]
