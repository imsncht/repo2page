# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod ./
# COPY go.sum ./ # Uncomment when dependencies are added
RUN go mod download

COPY . .
RUN go build -o /repo2page ./cmd/repo2page

# Final Stage
FROM alpine:latest

WORKDIR /
COPY --from=builder /repo2page /usr/local/bin/repo2page

# Basic runtime dependencies if needed (e.g., git, though we use tarballs mostly)
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["repo2page"]
CMD ["--help"]
