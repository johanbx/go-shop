FROM golang:1.22-alpine as builder

RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o server cmd/app/main.go

# note we use alpine and not scratch simply because we want to make
# outgoing https calls and therefor need ca-certificates.
# alpine is also better for debugging
FROM alpine:3
WORKDIR /app

RUN addgroup -g 1000 shop && adduser -u 1000 -D -S -G shop -s /bin/sh shop
RUN mkdir db && chown shop:shop db

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server server
COPY --from=builder /app/templates templates
COPY --from=builder /app/assets assets

USER shop:shop

ENTRYPOINT ["/app/server"]