FROM node:10-alpine AS builder

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . .

RUN npm run prod

FROM nginx:1.17.1-alpine

COPY --from=builder /app/dist/Student-API /usr/share/nginx/html