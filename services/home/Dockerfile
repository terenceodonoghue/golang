FROM golang:1.21 as build

WORKDIR /app

COPY go.mod go.sum *.go ./
COPY internal ./internal

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-home

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /docker-gs-home /docker-gs-home

EXPOSE 8080

CMD ["/docker-gs-home"]