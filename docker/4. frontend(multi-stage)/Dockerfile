#以node的alpine版本當鏡像as是用來額外命名，命名這個IMAGE叫builedr
FROM node:alpine as builder

#指定工作目錄
WORKDIR '/app'

#因為若COPY檔案跟之前的不同有任何變動的話就會執行RUN npm install
#所以不會事先用COPY將COPY . . 將所有的檔案載入，這樣會導致每次都要RUN npm install
COPY package.json . 

RUN npm install

COPY . .

#會執行package.json的build功能，若package.json沒寫則無功能
RUN npm run build 


FROM nginx

#沒有--from的話是抓當前資料夾，若有家則抓指定IMAGE的目錄
COPY --from=builder /app/build /usr/share/nginx/html
