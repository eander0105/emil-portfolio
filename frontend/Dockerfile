# Development stage
FROM node:18-alpine as dev

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5173

ENV HOST=0.0.0.0

CMD ["npm", "run", "dev"]

# Build stage
FROM node:18-alpine as build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build
