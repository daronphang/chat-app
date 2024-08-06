# Envoy

Envoy Proxy is an open-source edge and service proxy designed for cloud-native applications, and serves as a universal data plane for large-scale microservice service mesh architectures.

Envoy translates the HTTP/1.1 calls produced by the gRPC-web client into HTTP/2 calls that can be handled by those backend services (gRPC uses HTTP/2 for transport).

## Development

### Validation

```sh
$ envoy --mode validate -c my-envoy-config.yaml
```

### Run

```sh
$ envoy -c my-envoy-config.yaml
```
