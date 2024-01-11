FROM node:20 as frontend-builder

WORKDIR /app

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . .
RUN npm run build

FROM nginx:alpine

COPY --from=frontend-builder /app/build /usr/share/nginx/html
COPY run.sh .
RUN chmod +x run.sh

EXPOSE 80

ENV PUBLIC_BACKEND_URL=http://localhost:3000

CMD ["./run.sh"]
