FROM golang:1.21.5 as backend-builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# FROM alpine:latest  
FROM debian:stable-slim

VOLUME /data
WORKDIR /app

COPY --from=backend-builder /app/main .

EXPOSE 3000
ENV DB_PATH=/data/db.sqlite

CMD ["./main"]
