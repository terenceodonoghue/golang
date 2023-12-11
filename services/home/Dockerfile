FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum *.go ./
COPY internal ./internal

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /home-service

EXPOSE 8080

CMD ["/home-service"]