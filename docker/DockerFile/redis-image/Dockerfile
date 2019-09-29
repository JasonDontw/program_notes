#id 都是假設的

#use an existing docker image as a base
#指定一個鏡像，因為local沒有alpine image，所以到Docker Hub下載image，id為196d12cf6ab1
FROM alpine 

#download and install a dependency  安裝redis
#以196d12cf6ab1跑一個temporary container，container id為9054ef5db46a，執行apk add --update redis，完成後
#移除temporary container(Removing intermediate container 9054ef5db46a)，產生新的image 412c53a5addb
RUN apk add --update redis 

#tell the image what to do when it starts as a container
#以412c53a5addb跑一個temporary container，container id為174c749daf92，執行CMD產生一個以為start command的container，完成後
#移除temporary container(Removing intermediate container 174c749daf92)，產生新的image 00637dd534bf
CMD ["redis-server"]