# Development stage
FROM node:24-alpine AS dev

WORKDIR /app

COPY package*.json ./

RUN npm install

EXPOSE 5173

CMD ["npm", "run", "dev"]

# Build stage
FROM node:24-alpine AS build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# TODO: Run/production stage
