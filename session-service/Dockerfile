ARG WORKINGPATH=/app
ARG ENTRYPATH=/app
ARG CONFIG=development
ARG DEPLOYMENT_IMAGE=busybox:1.36.1

# Stage: BUILD
# Install dependencies first to maximize Docker layer caching.
FROM golang:1.22.2 AS build
ARG WORKINGPATH
ARG CONFIG
WORKDIR ${WORKINGPATH}

# Install packages.
COPY protobuf/ /protobuf
COPY session-service/go.mod session-service/go.sum ./
RUN go mod download

# Build from source code.
COPY session-service/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /http ./cmd/http

# Stage: DEPLOY
FROM $DEPLOYMENT_IMAGE
ARG WORKINGPATH
WORKDIR ${WORKINGPATH}

RUN wget "https://github.com/fullstorydev/grpcurl/releases/download/v1.9.1/grpcurl_1.9.1_linux_x86_64.tar.gz"; \
tar -zxvf grpcurl_1.9.1_linux_x86_64.tar.gz -C /bin

# When reading directory files in binary, path refers to
# the location of where the binary was built.
# To copy the files as reference.
COPY session-service/ ./

COPY --from=build /http /http
ENTRYPOINT ["/http"]