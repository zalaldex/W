# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files first for efficient caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Compile a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o bot ./cmd/bot

# Stage 2: Create a lightweight final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/bot .

EXPOSE 8080

CMD ["./bot"]
