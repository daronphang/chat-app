ARG WORKINGPATH=/app
ARG ENTRYPATH=/app
ARG CONFIG=development
ARG DEPLOYMENT_IMAGE=alpine:3.20

# Stage: BUILD
# Install dependencies first to maximize Docker layer caching.
FROM golang:1.22.2 AS build
ARG WORKINGPATH
ARG CONFIG
WORKDIR ${WORKINGPATH}

# Install packages.
COPY go.mod go.sum ./
RUN go mod download

# Build from source code.
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /http ./cmd/http

# Stage: DEPLOY
FROM $DEPLOYMENT_IMAGE
ARG WORKINGPATH
WORKDIR ${WORKINGPATH}

RUN apk add curl

# When reading directory files in binary, path refers to
# the location of where the binary was built.
# To copy the files as reference.
COPY . ./

COPY --from=build /http /http
ENTRYPOINT ["/http"]