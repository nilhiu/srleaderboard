FROM golang:1.23 AS builder

WORKDIR /usr/src/app

ENV CGO_ENABLED=0
ENV GOCACHE=/go-cache
ENV GOMODCACHE=/gomod-cache

COPY . .
RUN --mount=type=cache,target=/gomod-cache go mod download
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache go build -o srl ./cmd/main.go

FROM alpine:3.21

COPY --from=builder /usr/src/app/srl /bin/srl

ENTRYPOINT ["srl"]
