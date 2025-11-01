FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o velo ./cmd/velo

# final image
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/velo .
RUN mkdir -p /app/store

ENV VELO_ADDR=:8080 \
    VELO_CACHE_DIR=cache \
    VELO_MAX_IMAGE_SIZE=10485760 \
    VELO_WHITELISTED_DOMAINS=.* 

EXPOSE 8080
ENTRYPOINT ["./velo"]

