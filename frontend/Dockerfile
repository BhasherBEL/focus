FROM node:20 as frontend-builder

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build

FROM node:20-alpine

COPY --from=frontend-builder /app .

EXPOSE 4173

ENV PUBLIC_BACKEND_URL=http://localhost:3000

CMD ["npm", "run", "preview", "--", "--port", "4173", "--host", "0.0.0.0"]
