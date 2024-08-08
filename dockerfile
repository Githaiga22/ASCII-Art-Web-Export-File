# Build Stage
FROM golang:1.21-alpine3.18 AS builder
RUN apk add --no-cache bash
WORKDIR /app
COPY . .
RUN go build -o main ./main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY templates templates/
COPY static static/
COPY ascii-art/artstyles/ ascii-art/artstyles/
EXPOSE 8080
CMD ["./main"]docker