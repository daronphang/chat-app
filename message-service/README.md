# Message service

Responsibilities of chat service are as follows:

- Stores messages in NoSQL persistent store (Cassandra)
- Maintains channel-client and client-channel relationships
- If receiver is online, to forward to message to respective chat server
- If receiver is offline, to send push notification

Other design considerations to take note of:

- A group chat can be duplicated i.e. multiple groups having the same users inside
- 1-on-1 chats musts have a single channelId; for new chats, need to perform additional logic to check if the other user has already created
- For new chats created, need to broadcast to user's devices
- If two message sent in sequence to a user, and the second message arrives first, there is a mismatch in ordering; this can be resolved with prevMsgId where if an inconsistency is found, a history catch-up can be requested

## Development

### Cassandra

1. Spin up Docker instance

```sh
$ docker run --rm --name cassandra -p 9042:9042 -d cassandra:5.0
```

### RabbitMQ

1. Spin up Docker instance

```sh
$ docker run --rm --name rabbitmq -p 5672:5672 -p 15672:15672 -d rabbitmq:3.13-management
```

### Kafka

1. Spin up Docker instance

```sh
$ docker run --rm --name kafka -p 9092:9092 -d apache/kafka:3.7.0
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
