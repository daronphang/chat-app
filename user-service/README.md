# User service

## First time setup

### etcd

1. Spin up Docker instance. Host IP is the IP of docker container

```sh
$ docker run -d -p 4001:4001 -p 2380:2380 -p 2379:2379 \
 --name etcd quay.io/coreos/etcd:v3.5.15 \
  /usr/local/bin/etcd \
 --name etcd0 \
 --advertise-client-urls http://HOST_IP:2379,http://HOST_IP:4001 \
 --listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
 --initial-advertise-peer-urls http://HOST_IP:2380 \
 --listen-peer-urls http://0.0.0.0:2380 \
 --initial-cluster-token etcd-cluster-1 \
 --initial-cluster etcd0=http://HOST_IP:2380 \
 --initial-cluster-state new
```

### Kafka

1. Spin up Docker instance

```sh
$ docker run --rm --name kafka -p 9092:9092 -d apache/kafka:3.7.0
```

### PostgreSQL

1. Spin up Docker instance

```sh
$ docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=password postgres:16.4
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

## Deployment

### Docker

As protobuf files are stored in the parent directory, Docker's build context needs to be specified from there.

```sh
$ cd path/to/parent/directory
$ docker build -t user-service -f ./user-service/Dockerfile .
```

### Docker compose

```sh
$ docker network create -d bridge chatapp
$ docker compose -f docker-compose-staging.yaml up -d
```
