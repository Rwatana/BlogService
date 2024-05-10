FROM node:14

ENV CI=true
ENV WDS_SOCKET_PORT=0

WORKDIR /app
COPY package.json ./
RUN npm install

COPY ./ ./

CMD ["npm", "start"]

# # test sentence
# FROM node:alpine

# ENV CI=true
# ENV WDS_SOCKET_PORT=0

# WORKDIR /app
# COPY package.json ./
# COPY ./ ./

# CMD ["tail", "-f", "/dev/null"]