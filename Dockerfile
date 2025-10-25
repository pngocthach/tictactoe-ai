# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/main .


# Stage 2: Final
FROM alpine:latest

COPY --from=builder /app/main /main
COPY --from=builder /app/web /web

EXPOSE 8080

ENTRYPOINT ["/main"]