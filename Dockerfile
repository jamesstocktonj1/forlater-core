# syntax=docker/dockerfile:1

# build container
FROM golang:1.21 AS builder

WORKDIR /app

COPY . ./
RUN go mod download

ARG CGO_ENABLED=0
RUN go build -ldflags="-w -s" -o ./proxy cmd/proxy/main.go

# run container
FROM scratch

WORKDIR /app

COPY --from=builder /app/proxy ./proxy

EXPOSE 8000
CMD ["./proxy"]