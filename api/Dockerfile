FROM node:alpine

WORKDIR /usr/src/app

# Copy package.json and lock file(s)
COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000
CMD [ "npm", "start" ]