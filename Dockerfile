FROM golang:1.22-alpine AS build
WORKDIR /src

# Copy module files and tidy, then copy the rest and build
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN apk add --no-cache git
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bot ./cmd/monospace-bot

FROM alpine:3.20
RUN apk add --no-cache ca-certificates

# Create /data for persistent storage and runtime files
RUN mkdir -p /data
VOLUME ["/data"]

COPY --from=build /bot /bot
EXPOSE 8080
ENTRYPOINT ["/bot"]
