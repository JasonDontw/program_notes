#以node的alpine版本當鏡像
FROM node:alpine 

#指定工作目錄
WORKDIR /usr/app

COPY ./ ./
#第一個./是HOST的根目錄第2個./是image的根目錄
#COPY將整個HOST根目錄複製到image裡面去
RUN npm install

CMD ["npm","start"]