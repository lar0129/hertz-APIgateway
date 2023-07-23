# hertz-APIgateway

## 1.项目结构

```shell
├── LICENSE
├── README.md

├── hertz_gateway
│   ├── biz
│   │   ├── clientprovider ## kitex client provider
│   │   │   └── clientprovider.go
│   │   ├── handler ## http Api layer
│   │   │   ├── demo ## 测试用
│   │   │   │   └── student_service.go
│   │   │   ├── gateway ## idl manager api
│   │   │   │   └── idl_service.go
│   │   │   ├── ping.go ## 测试用
│   │   │   ├── service_router.go ## router泛化调用 api
│   │   │   └── teacher ## 测试用
│   │   │       └── teacher_service.go
│   │   ├── idl ## idl相关处理
│   │   │   ├── idl_management.go ## idl manager 后端
│   │   │   └── idl_provider.go ## idl provider
│   │   ├── model ## hertz自动生成
│   │   │   ├── demo
│   │   │   │   └── stu.go
│   │   │   ├── gateway
│   │   │   │   └── gateway.go
│   │   │   └── teacher
│   │   │       └── teacher.go
│   │   └── router ## 路由解析
│   │       ├── demo
│   │       │   ├── middleware.go
│   │       │   └── stu.go
│   │       ├── gateway
│   │       │   ├── gateway.go
│   │       │   └── middleware.go
│   │       ├── register.go
│   │       └── teacher
│   │           ├── middleware.go
│   │           └── teacher.go
│   ├── build.sh
│   ├── go.mod
│   ├── go.sum
│   ├── hertz.demo
│   ├── main.go
│   ├── router.go ## 路由解析
│   ├── router_gen.go
│   └── script
│       └── bootstrap.sh


├── idl ## thrift IDL 文件
│   ├── gateway.thrift
│   ├── stu.thrift
│   └── teacher.thrift

├──student-service ## 测试用微服务
│  ├──...
├──teacher-service  ## 测试用微服务
│  ├──...
```

![image-20230722122427181](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230722122427181.png)

## 2.网关接口约定

![image-20230723120718797](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230723120718797.png)

* HTTP约定 : IP address/gateway/ServiceName/MethodName
  * ServiceName为在etcd注册中心注册使用的名字
  * MethodName必须与微服务中的方法名大小写一致
  * 目前只支持POST类型请求

## 3.本地部署步骤

1.打开etcd注册中心

* 运行etcd --log-level debug

2.打开服务端

* 从根目录打开student_service文件夹（或teacher_service）
* 运行命令
  * sh build.sh
  * sh output/bootstrap.sh
  * 显示etcd注册成功即可

3.打开api网关服务

* 从根目录打开hertz_gateway文件夹。
* 运行命令
  * go build .
* 运行生成的可执行文件./hertz.demo

4.用postman或curl命令进行接口测试

* 接口使用示例
  * ![image-20230723120439639](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230723120439639.png)