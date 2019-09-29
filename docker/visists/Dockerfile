#以node的alpine版本當鏡像
FROM node:alpine 

#指定工作目錄
WORKDIR /app

COPY package.json .

RUN npm install
COPY . .

CMD ["npm","start"]