FROM golang:1.12.5-alpine AS builder

ADD . /root/app

RUN apk add --no-cache git

# Download modules
RUN cd /root/app && \
    GO111MODULE=on GOPROXY=https://gocenter.io go mod download

# Build microservices
RUN cd /root/app && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM debian:buster
COPY --from=builder /root/app/security-sample-app /bin

# Set runtime user to non-root
USER 1000

CMD ["security-sample-app"]
