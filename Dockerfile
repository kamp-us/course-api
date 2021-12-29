# Building the binary of the App
FROM golang:latest AS builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/course-api /app/cmd/course-api/main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine
WORKDIR /app
RUN apk add --no-cache nano git curl
# COPY --from=builder /feedback-api/internal/configs/dev.env .env
COPY --from=builder bin/course-api course-api

CMD ["./course-api"]

EXPOSE 80
