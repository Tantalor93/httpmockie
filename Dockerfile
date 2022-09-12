# syntax = docker/dockerfile:1.2
# Builder image
FROM golang:1.19.1-alpine3.15 as builder

RUN apk add --update build-base libmnl-dev

RUN mkdir -p /httpmockie
WORKDIR /httpmockie

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOMODCACHE=/go/pkg/mod \
    GOCACHE=/root/.cache/go-build \
    make build

# Runtime image
FROM alpine:3.15

RUN apk --no-cache add ca-certificates
COPY --from=builder /httpmockie/bin/httpmockie /app/httpmockie
WORKDIR /app

ENTRYPOINT ["./httpmockie"]
