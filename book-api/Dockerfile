FROM node:20-alpine

WORKDIR /app

COPY package.json database.json openapi.yaml ./

RUN npm install

EXPOSE 3000 4010 8080

CMD ["npm", "start"]
