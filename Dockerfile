FROM golang:1.16.4 as builder

COPY . /app
WORKDIR /app
RUN make build

# Runtime image
FROM alpine:3.13

COPY --from=builder /app/mhz19b /mhz19b
ENTRYPOINT ["/mhz19b"]