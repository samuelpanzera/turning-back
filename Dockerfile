FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

RUN adduser -D -g '' appuser

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /build/main /app/main

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/main"]