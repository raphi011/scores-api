FROM node:latest

WORKDIR /scores/web
COPY ./web/package*.json ./
RUN npm install

WORKDIR /scores
COPY . .

WORKDIR /scores/web
RUN npm run build

EXPOSE 3000 

ENTRYPOINT ["npm"]
CMD ["start"]