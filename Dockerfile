# use the official Bun image
# see all versions at https://hub.docker.com/r/oven/bun/tags
FROM imbios/bun-node AS frontend
WORKDIR /usr/src/app

COPY . .

RUN bun i
RUN bun run build

FROM golang:1.22 AS builder

ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src

COPY --from=frontend /usr/src/app/ .

# Build Go binary
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go mod download 
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,mode=0755,target=/go/pkg/mod go build -o serve ./cmd/serve

FROM alpine:latest

WORKDIR /app

COPY --from=builder /go/src/serve /

RUN apk add --no-cache \
    unzip \
    tzdata \
    ca-certificates

EXPOSE 8080

CMD ["/serve"]
