FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bot ./cmd/bot

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
RUN mkdir -p /data
COPY --from=build /bot /bot
EXPOSE 8080
ENTRYPOINT ["/bot"]