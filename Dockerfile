FROM node:carbon
# FROM golang:1.9


WORKDIR /frontend
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

EXPOSE 3000 

CMD ["npm", "start"]
# CMD["./backened/scores-backend"]