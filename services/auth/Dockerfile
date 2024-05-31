FROM golang:1.21 as build

WORKDIR /app

COPY go.mod go.sum *.go ./
COPY internal ./internal

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-auth

FROM scratch
COPY --from=build /docker-gs-auth /docker-gs-auth

EXPOSE 8080

CMD ["/docker-gs-auth"]