# Message service

## First time setup

### Cassandra

1. Spin up Docker instance

```sh
$ docker run --rm --name cassandra -p 9042:9042 -d cassandra:5.0
```

### Kafka

1. Spin up Docker instance

```sh
$ docker run --rm --name kafka -p 9092:9092 -d apache/kafka:3.7.0
```

## Development

### Wire (DI)

https://github.com/google/wire

1. Write Providers and Wire functions

2. Generate Wire code

```sh
$ cd path/to/root/directory
$ wire ./internal
$ go generate ./internal # once wire_gen.go is created, can regenerate using this
```

### HTTP server

1. Run server

```sh
$ cd path/to/root/directory
$ go run cmd/http/main.go
```

### Kafka workers

1. Run workers

```sh
$ cd path/to/root/directory
$ go run cmd/worker/main.go
```

## Testing

Before running tests, set environment variable GO_ENV to 'TESTING'.

```sh
$ export GO_ENV=TESTING
```

### Mockery

To generate mocks from interfaces, use Mockery.

https://github.com/vektra/mockery

```sh
$ brew install mockery
$ cd path/to/root/directory
$ mockery # reads from .mockery.yaml config file
```

### Running unit tests

```sh
$ cd path/to/root/directory
$ go test ./... -v
$ go test ./... -v -coverpkg=./...
```

### Running integration tests

1. Start containers with Docker Compose

```sh
$ cd path/to/root/directory
$ docker compose -f docker-compose-testing.yaml up -d
```

## Deployment

### Docker

As protobuf files are stored in the parent directory, Docker's build context needs to be specified from there.

```sh
$ cd path/to/parent/directory
$ docker build -t message-service -f ./message-service/Dockerfile .
$ docker run -d -p 50051:50051 message-service
```

### Docker compose

```sh
$ docker network create -d bridge chatapp
$ docker compose -f docker-compose-staging.yaml up -d
```
