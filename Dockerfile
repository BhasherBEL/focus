#Frontend
FROM node:20 as frontend-builder

WORKDIR /app

COPY frontend/package.json .
COPY frontend/package-lock.json .
RUN npm install

COPY frontend .
RUN npm run build

# Backend
FROM golang:1.21.5 as backend-builder
WORKDIR /app
COPY backend/. .
RUN CGO_ENABLED=1 GOOS=linux go build -o main

# Combined
FROM debian:stable-slim
WORKDIR /app

RUN apt-get update && \
    apt-get install -y nginx

COPY --from=frontend-builder /app/build /var/www/html

COPY --from=backend-builder /app/main /app/backend/main

COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh

VOLUME /data

EXPOSE 80 3000

ENV DB_PATH=/data/db.sqlite
ENV PUBLIC_BACKEND_URL=http://localhost:3000


CMD ["/app/run.sh"]
