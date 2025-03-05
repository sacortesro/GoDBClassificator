FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

EXPOSE 8080

RUN go build -o main main.go

CMD ["./main"]