FROM golang:1.14 as build-env

RUN mkdir /app
WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

#### Building the binary
RUN go build -o run *.go

### Build final minimal image
FROM golang:1.14
COPY --from=build-env /app/run /app/run
WORKDIR /app
ENTRYPOINT ["./run"]
