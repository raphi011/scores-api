FROM node:carbon

WORKDIR /web
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

EXPOSE 3000 

ENTRYPOINT ["npm"]
CMD ["start"]