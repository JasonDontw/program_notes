#以node的alpine版本當鏡像
FROM node:alpine 

#指定工作目錄
WORKDIR '/app'

#因為若COPY檔案跟之前的不同有任何變動的話就會執行RUN npm install
#所以不會事先用COPY將COPY . . 將所有的檔案載入，這樣會導致每次都要RUN npm install
COPY package.json . 

RUN npm install

COPY . .

CMD ["npm","start"]