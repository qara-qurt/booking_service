# Multi-stage building
# 1. Build the compiled binary
FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app ./cmd/app

# 2. Build the final image and run binary file
FROM alpine
WORKDIR /build
COPY --from=builder /build/app /build/app
COPY .env /build/.env
COPY internal/repository/postgres/migrations /build/migrations

# Install migrate CLI
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin

# Set up environment for testing
ENV DATABASE_URL=postgres://postgres:password@postgres:5432/postgres?sslmode=disable

# Expose the port
EXPOSE 3000

# Run migration and start the app
CMD ["sh", "-c", "migrate -path /build/migrations -database ${DATABASE_URL} up && /build/app"]
