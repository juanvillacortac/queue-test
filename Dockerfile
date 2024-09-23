# use the official Bun image
# see all versions at https://hub.docker.com/r/oven/bun/tags
FROM imbios/bun-node AS frontend
WORKDIR /usr/src/app

COPY . .

RUN bun i
RUN bun run build

FROM golang:1.22-alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src

# Install dependencies
RUN apk --update --no-cache add ca-certificates

COPY --from=frontend /usr/src/app/ .

# Build Go binary
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod go mod download 
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,mode=0755,target=/go/pkg/mod go build -o serve ./cmd/serve

FROM alpine:latest

COPY --from=builder /go/src/serve /app/

WORKDIR /app

RUN apk add --no-cache \
    unzip \
    ca-certificates \
    # this is needed only if you want to use scp to copy later your pb_data locally
    openssh

EXPOSE 8080

CMD ["./serve"]
