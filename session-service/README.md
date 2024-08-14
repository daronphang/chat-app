# Session service

Responsibilities of session service are as follows:

- Maintain heartbeat session with users and storing chat server metadata for each user
- Perform broadcasting events i.e. if user is online, perform A, else, perform B

## Development

### Redis

1. Spin up Docker instance

```sh
$ docker run --rm --name redis -d -p 6379:6379 redis
```

### Wire (DI)

https://github.com/google/wire

1. Write Providers and Wire functions

2. Generate Wire code

```sh
$ cd path/to/root/directory
$ wire ./internal
$ go generate ./internal # once wire_gen.go is created, can regenerate using this
```

### Web server

1. Run server

```sh
$ cd path/to/root/directory
$ go run cmd/rest/main.go
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
