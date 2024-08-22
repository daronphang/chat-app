# Protobuf

# Setup

1. Install protobuf compiler

https://grpc.io/docs/protoc-installation/

2. Install compiler plugins for Go, Javascript and gRPC-web

```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$ sudo npm install -g protoc-gen-js

$ sudo npm install -g protoc-gen-grpc-web
```

3. Update PATH so that the protoc compiler can find the plugins

```sh
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Usage

1. Update the respective .proto files

2. Run protobuf compiler for Go modules and gRPC-web client (React). Make sure dependencies are also compiled e.g. common

```sh
$ cd path/to/root/directory
$ protoc --proto_path=. \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --js_out=import_style=commonjs:../chat-ui/src \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:../chat-ui/src \
    ./proto/SERVICE_NAME/NAME.proto
```

3. Import files into service
