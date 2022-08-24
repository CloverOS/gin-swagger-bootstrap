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
