# go-zero-mall

参考go-zero官方公众号测试编写

## docker启动相关服务

docker-compose up -d

##  redis工具写入对应的key
hset rpc.auth.user userapi  xxx

hset rpc.auth.product productapi  xxx

hset rpc.auth.order orderapi  xxx

hset rpc.auth.pay payapi  xxx
