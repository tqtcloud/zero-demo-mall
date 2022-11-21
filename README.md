# go-zero 实践

redis 实现rpc认证
```shell
# 设置 hset key值
 HSET rpc:auth:user userapi  6jKNZbEpYGeUMAifz10gOnmoty3TV
 
 # 查询
 HGET rpc:auth:user userapi
```
