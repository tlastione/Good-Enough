FROM golang:1.25.4-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o myapp/ ./... 

FROM alpine:latest

WORKDIR /app

RUN addgroup -S appuser \
 && adduser -S -G appuser -H -s /sbin/nologin appuser

COPY --from=builder --chown=appuser:appuser /app/myapp/cmd /app/myapp

USER appuser

ENTRYPOINT ["/app/myapp"]
