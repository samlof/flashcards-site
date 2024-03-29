# syntax=docker/dockerfile:1.0.0-experimental

FROM golang:1.14.6-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . .

FROM base AS build
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux GOARCH=amd64 go build -o /out/app .

FROM base AS test
ENV CGO_ENABLED=1
RUN apk add build-base
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    go test -v -race ./...


# FROM golangci/golangci-lint:v1.30-alpine AS lint-base

# FROM base AS lint
# RUN --mount=target=. \
#    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
#    --mount=type=cache,target=/root/.cache/go-build \
#    --mount=type=cache,target=/root/.cache/golangci-lint \
#    golangci-lint run --timeout 10m0s ./...

FROM scratch AS bin
COPY --from=build /out/app /