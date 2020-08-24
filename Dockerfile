# Builder image
FROM golang:1.14 AS builder

RUN mkdir -p /go-http-mock
WORKDIR /go-http-mock

COPY . .

RUN make build

# Runtime image
FROM alpine:3.9

RUN apk --no-cache add ca-certificates
COPY --from=builder /go-http-mock/bin/go-http-mock /app/go-http-mock
WORKDIR /app

ENTRYPOINT ["./go-http-mock"]
