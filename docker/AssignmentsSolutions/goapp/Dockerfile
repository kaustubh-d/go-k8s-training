FROM golang:1.16-alpine as builder
WORKDIR /app
COPY . .

RUN go build && \
    chmod 777 hello

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/hello /usr/bin/hello

ENTRYPOINT ["/usr/bin/hello"]
