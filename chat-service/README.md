# Chat service

## First time setup

### etcd

```sh
$ go get go.etcd.io/etcd/etcdctl/v3
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

### IP address

```sh
$ hostname -i # private
$ wget -qO- ifconfig.me # public
$ docker inspect \
  -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' container_name_or_id
```

### Docker

```sh
$ cd path/to/root/directory
$ docker build -t chat-service .
```

### Docker compose

```sh
$ docker network create -d bridge chatapp
$ docker compose -f docker-compose-staging.yaml up -d
```
