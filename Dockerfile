FROM golang:1.16.4 as builder

ARG GOOS
ARG GOARCH

COPY . /app
WORKDIR /app
RUN make build -e GOOS=$GOOS GOARCH=$GOARCH

# Runtime image
FROM alpine:3.13

COPY --from=builder /app/mhz19b /mhz19b
ENTRYPOINT ["/mhz19b"]
