FROM node:carbon

ARG backend_url
ENV BACKEND_URL $backend_url

WORKDIR /frontend
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

EXPOSE 3000 

CMD ["npm", "start"]