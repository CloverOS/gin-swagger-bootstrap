# gin-swagger-bootstrap
swagger前端页面ui增强

使用
```sh
go get -u https://github.com/CloverOS/gin-swagger-bootstrap
```

在gin路由配置中配置如下
```go
r.GET("/swagger/*any", ginSwaggerBootstrap.WrapHandler())
```

该项目仅做前端ui的增强替换，不做其他修改

具体的文档请参考 [gin-swagger](https://github.com/swaggo/gin-swagger)

预览图
![Image text](https://raw.githubusercontent.com/CloverOS/gin-swagger-bootstrap/main/preview/1.jpg)
![Image text](https://raw.githubusercontent.com/CloverOS/gin-swagger-bootstrap/main/preview/2.jpg)
![Image text](https://raw.githubusercontent.com/CloverOS/gin-swagger-bootstrap/main/preview/3.png)
