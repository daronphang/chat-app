# Protobuf

# Setup

1. Install protobuf compiler

https://grpc.io/docs/protoc-installation/

2. Install compiler plugins for Go

```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

3. Update PATH so that the protoc compiler can find the plugins

```sh
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Usage

1. Update the respective .proto files

2. Run protobuf compiler

```sh
$ cd path/to/root/directory
$ protoc --proto_path=. \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./SERVICE_NAME/NAME.proto
```

3. Import files into service
