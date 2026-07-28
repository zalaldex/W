FROM golang:1.22-alpine AS build
WORKDIR /src

# Copy all repository files
COPY . .

# Download dependencies and sync go.mod / go.sum
RUN go mod tidy

# Build the binary
RUN CGO_ENABLED=0 go build -o /bot ./cmd/bot

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
RUN mkdir -p /data
COPY --from=build /bot /bot
EXPOSE 8080
ENTRYPOINT ["/bot"]
