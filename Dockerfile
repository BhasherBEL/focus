FROM node:20 as frontend-builder
WORKDIR /app
COPY frontend/. .
RUN npm install
RUN npm run build

FROM golang:1.21.5 as backend-builder
WORKDIR /app
COPY backend/. .
RUN CGO_ENABLED=1 GOOS=linux go build -o main

FROM debian:stable-slim
WORKDIR /app

COPY --from=frontend-builder /app /app/frontend

COPY --from=backend-builder /app/main /app/backend/main

RUN apt-get update && apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_20.x | bash -
RUN apt-get install -y nodejs

VOLUME /data

EXPOSE 4173 3000

ENV DB_PATH=/data/db.sqlite
ENV PUBLIC_BACKEND_URL=http://localhost:3000

COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh

CMD ["/app/run.sh"]
