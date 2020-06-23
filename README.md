### swag文档
```
swag init


Param Type
    - query
    - path
    - header
    - body
    - formData
Data Type
    - string (string)
    - integer (int, uint, uint32, uint64)
    - number (float32)
    - boolean (bool)
    - user defined struct
```

### 构建
```shell script
go build -o bastion main.go

./bastion


ps -ef | grep bastion 

kill -9 


docker run -d \
    -p 6379:6379 \
    -v /root/data/redis.conf:/usr/local/etc/redis/redis.conf \
    --privileged=true \
    --name myredis\
     redis \
    redis-server /usr/local/etc/redis/redis.conf

```

```
必要拷贝文件夹
public
conf
templates
```


关于signal 
 
查看博客 [Linux Signal及Golang中的信号处理][https://colobu.com/2015/10/09/Linux-Signals/]
```shell script
# 优雅关闭
kill -15

# 强制关闭
kill -9
```


### 自动生成model
```shell script
https://github.com/xxjwxc/gormt

gormt

记得将主键露出


# 全部 json:"-" 替换 json:"id"

```



### 分页问题

https://github.com/jinzhu/gorm/issues/1752#issuecomment-454457879

```
 db = db.Model(&model.ParsedData{}).
        Where(where).
        Order("created_at DESC").
        Count(&page.Total).
        Limit(page.Limit). // Do limit and offset after count, but you now need Model before.
        Offset(page.Offset).
        Find(&parsedData)
```