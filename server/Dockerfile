FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod tidy

COPY . .

RUN make docs && make build 

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/veritas .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./veritas"]
