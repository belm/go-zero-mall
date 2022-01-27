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

## Jaeger 链路追踪
![image](https://user-images.githubusercontent.com/551218/151284144-7e6a6711-e6f7-4c9e-a100-e9a2fbf8f692.png)


## 感谢
[go-zero](https://go-zero.dev/)

[微服务实践公众号](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2ODU1MTI0OA==&action=getalbum&album_id=2085775054620917763&scene=173&from_msgid=2247484980&from_itemidx=1&count=3&nolastread=1#wechat_redirect)
