# go-zero-mall

参考go-zero官方公众号测试编写

## docker启动相关服务

docker-compose up -d

##  redis工具写入对应的key
hset rpc.auth.user userapi  xxx

hset rpc.auth.product productapi  xxx

hset rpc.auth.order orderapi  xxx

hset rpc.auth.pay payapi  xxx

## prometheus
![image](https://user-images.githubusercontent.com/551218/151104925-9534687b-7d4a-43d4-9894-08404b6dbc70.png)

选择 Graph 菜单，在查询输入框中输入 {path="api接口地址"} 或者 {method="rpc接口方法"} 指令，即可查看监控指标。

![image](https://user-images.githubusercontent.com/551218/151105060-973a5682-366e-4891-bc92-f17287a4f07f.png)


## grafana监控图
文件go-zero-mall-grafana.json是grafana模板，直接导入即可

![image](https://user-images.githubusercontent.com/551218/151104851-760e5c31-9a5a-47cc-857c-ecac58852c7b.png)
