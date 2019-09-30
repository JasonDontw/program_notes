const express = require('express');
const redis = require('redis');
const process = require('process')
const app = express();
const client = redis.createClient({
  host: 'redis-server', //會參考到docker-compose.yml
  port: 6379 //創建到port:6379
});//連接到服務器

client.set('visits',0);//設定鍵值為visits為0

app.get('/', (req, res) => {
  client.get('visits',(err,visits)=>{//取得鍵值為visits得值
    res.send('Number of visits is'+ visits);
    client.set('visits',parseInt(visits)+1);//設定鍵值為visits得值
  }) 
});


app.listen(8080, () => {
  console.log('Listening on port 8080');
});