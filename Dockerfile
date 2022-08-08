# Builder image
FROM golang:1.14 AS builder

RUN mkdir -p /httpmockie
WORKDIR /httpmockie

COPY . .

RUN make build

# Runtime image
FROM alpine:3.9

RUN apk --no-cache add ca-certificates
COPY --from=builder /httpmockie/bin/httpmockie /app/httpmockie
WORKDIR /app

ENTRYPOINT ["./httpmockie"]
