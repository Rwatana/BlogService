FROM node:18.20.0-alpine

ENV CI=true
ENV WDS_SOCKET_PORT=0

WORKDIR /app
COPY package.json ./
COPY ./ ./

CMD ["npm", "start"]

