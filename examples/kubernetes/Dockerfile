FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o producer ./cmd/producer

FROM scratch
COPY --from=builder /app/producer /app/producer
CMD ["/app/producer"]