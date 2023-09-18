# Base image létrehozása

FROM golang:1.16 AS builder

WORKDIR /app
COPY app/go.mod ./
RUN go mod download
COPY app/ ./
RUN go build -o myapp

# Alkalmazás futtató image létrehozása
FROM alpine:latest
WORKDIR /app
RUN apk add libc6-compat
COPY --from=builder /app/myapp .
CMD ["./myapp"]