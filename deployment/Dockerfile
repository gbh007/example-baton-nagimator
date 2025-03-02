FROM golang:1.23.6 AS builder

RUN useradd -u 1000 nonroot
WORKDIR /build
COPY . .

RUN --mount=type=cache,target="$(go env GOMODCACHE)" \
    --mount=type=cache,target="$(go env GOCACHE)" \
    CGO_ENABLED=0 go build \
    -trimpath \
    -o /build/main \
    cmd/main.go

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/timezone /etc/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /build/main /app/main

COPY ./internal/controller/index.html /app/internal/controller/index.html

USER nonroot
WORKDIR /app
ENTRYPOINT ["/app/main"]
CMD ["-d", "-conn=/app/data/main.db"]