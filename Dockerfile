FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /bot .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=build /bot /bot
EXPOSE 8080
ENTRYPOINT ["/bot"]
