# hertz-APIgateway

[TOC]

## 1.概述

### 1.1. 简介

* Gateway项目：https://github.com/lar0129/hertz-APIgateway
* IDL-Management-Platform项目前端：https://github.com/wangchunchia/cwg-frontend
  * IDL-Management-Platform后端为方便部署测试，目前与gateway耦合
* 视频演示：https://lar0129-video.oss-cn-nanjing.aliyuncs.com/7%E6%9C%8831%E6%97%A5.mp4

​		在微服务架构中，大型服务都被拆分成了独立的微服务，每个微服务通常会以RESTFUL API的形式对外提供服务。但是对客户端而言，我们可能需要在一个页面上显示来自不同微服务的数据，此时就会需要一个统一的入口来进行API的调用。**API Gateway就在此场景下充当了多个服务的大门，系统的统一入口**

​		在本项目中，API 网关作为系统的入口点，接收所有传入的 HTTP 请求，解析请求中的url，以JSON格式将它们转发到相关后端RPC服务器，检索响应，然后将其返回给客户端。



**技术选型**

* 编程语言：golang >= 1.20
* API-Gateway服务：hertz框架、Kitex Client
* RPC服务：Kitex Service框架
* Registry Center：etcd
* RPC 协议：thrift



**优点**

* 为大型分布式系统提供了可扩展且可靠的解决方案，提供一个统一的对外接口
* 采用API Gateway可以与微服务注册中心连接，实现微服务无感知动态扩容。
* API Gateway对于无法访问的服务，可以做到自动熔断，无需人工参与。
* API Gateway做为系统统一入口，我们可以将各个微服务公共功能放在API Gateway中实现，以尽可能减少各服务的职责。
* API Gateway帮助实现客户端的负载均衡策略。



### 1.2. 项目结构

```shell
├── hertz_gateway
│   ├── biz
│   │   ├── clientprovider		  ## kitex client provider
│   │   ├── handler 		 	  ## http Api layer
│   │   │   ├── gateway 	 	  		## idl manager api
│   │   │   ├── service_router.go 		## router泛化调用 api
│   │   │   ├── ping.go 		  		## 测试网关是否成功运行 api
│   │   ├── idl 				  ## idl相关处理
│   │   │   ├── idl_management.go 		## idl manager 后端
│   │   │   └── idl_provider.go   		## idl provider
│   │   ├── model 				  ## hertz自动生成
│   │   └── router 				  ## 路由解析
│   ├── main.go
│   ├── router.go 				  ## 解析服务名与方法名
│   ├── router_gen.go

├── test ## 测试相关
│   ├── output  		## jmeter测试结果
│   └── cloudwego.jmx 	## jmeter测试配置文件

├── idl ## thrift IDL 文件
│   ├── gateway.thrift  ## 该网关的  IDL 文件
│   ├── stu.thrift  	## RPC服务端 IDL 文件
│   └── teacher.thrift 	## RPC服务端 IDL 文件

├──student-service 		## 测试用RPC服务
│  ├──...
├──teacher-service  	## 测试用RPC服务
│  ├──...
```

* 该网关分为5个模块
  * API Layer层（biz/handler包）：接受HTTP的POST请求，并交由Routing层解析
    * 实现了对任意origin的**跨域**，可按需求在main.go更改
  * Routing层（biz/router包与./router.go）：从url中解析出转发到的服务名和方法名，并决定发往何处
    * **负载均衡**目前采用随机权重策略。可按需求更改
  * Kitex Client Provider层（biz/clientprovider包）：根据路由结果，获取泛化调用客户端
    * 通过etcd Registry Center实现**服务发现**
    * 通过IDL Provider 解析IDL文件
    * 采用**json泛化调用**，网关扩展性强
  * IDL Provider层（biz/idl/idl_provider.go文件）
    * 往 IDL Managerment 发送请求，获取服务名对应IDL文件并进行解析，返回解析后的IDL对象。
    * 当服务端的IDL文件有变更时，IDL Provider层的go线程**自动实现IDL-update**
  * IDL Managerment层（biz/idl/idl_management.go文件）
    * 存储 [服务端名称 -- 对应IDL路径 -- IDL文件描述] 的三者映射
    * 提供对IDL路径的增删改查功能
      * 演示视频：https://lar0129-video.oss-cn-nanjing.aliyuncs.com/7%E6%9C%8831%E6%97%A5.mp4
    * 当新增RPC服务端时，只需往IDL-managerment平台**添加对应服务名与idl文件路径**，即可实现网关注册
    * 目前只支持**thrift协议**的IDL文件

![image-20230722122427181](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230722122427181.png)







## 2.快速开始

### 2.1. 环境配置（以Ubuntu/Debian为例）

1.  安装etcd:

```shell
## 下载https://github.com/etcd-io/etcd/releases
curl -LO https://github.com/etcd-io/etcd/releases/download/v3.5.9/etcd-v3.5.9-linux-amd64.tar.gz

## 解压
tar xvf etcd-v3.5.9-linux-amd64.tar.gz

## 复制到 /usr/local/bin
cd etcd-v3.5.9-linux-amd64.tar.gz
sudo cp etcd /usr/local/bin/
sudo cp etcdctl /usr/local/bin/
```

2. golang:

```sh
## 安装golang
curl -LO https://golang.google.cn/dl/go1.20.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz

## 修改环境变量
sudo vim /etc/profile
export PATH=$PATH:/usr/local/go/bin
source /etc/profile

## 验证环境
go version
```

3. hz安装

```
go install github.com/cloudwego/hertz/cmd/hz@latest
```

4. kitex安装

```
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```

5. thriftgo安装

```
go install github.com/cloudwego/thriftgo@latest
```



### 2.2.  简单部署

1. 打开etcd注册中心

```sh
etcd --log-level debug
## 确保etcd注册在2379端口
```



2. 打开api网关服务

* 从根目录打开hertz_gateway文件夹。

```shell
## 构建可执行文件
go build .

## 运行可执行文件
./hertz.demo

## 默认注册在8888端口
```



3. 打开任意服务端（RPC服务端需在etcd服务注册）	

* 从根目录打开teacher_service文件夹（或student_service）

```shell
## 构建可执行文件
sh build.sh

## 运行可执行文件
sh output/bootstrap.sh

## 确认在etcd注册成功。否则网关无法实现服务发现
```



4. 用idl-management-platform前端平台或postman进行接口测试

* idl-management-platform前端平台测试示例
  * 演示视频：https://lar0129-video.oss-cn-nanjing.aliyuncs.com/7%E6%9C%8831%E6%97%A5.mp4

* postman测试示例
  * 先往IDL管理平台添加**任意服务**名称及其对应的IDL文件路径（IDL目前为本地文件）
  * ![image-20230723122836025](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230723122836025.png)
  * 再通过网关调用相应server的方法接口
  * ![image-20230723120439639](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230723120439639.png)

## 3.特性

### 3.1. 接口约定

![image-20230723120718797](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230723120718797.png)

* HTTP约定 : POST：\<IP address>/gateway/ServiceName/MethodName
  * ServiceName为在etcd注册中心注册使用的名字
  * MethodName必须与微服务中的方法名大小写一致
  * 目前只支持POST类型请求

* IDL管理平台接口约定（目前只支持POST类型请求）：

  * 增：\<IP address>/add-service

    * Body例：

      ```json
      { 
       "serviceName" : "teacher-server",
       "serviceIdl" : "../idl/teacher.thrift",
       "serviceIdlContent" : "teacher's IDL" 
      }
      ```

  * 删：\<IP address>/delete-service

    * Body例：

      ```json
      { "serviceName" : "teacher-server"}
      ```

  * 改：\<IP address>/update-service

    * Body例：

      ```json
      { 
          "serviceName" : "teacher-server",
      	"serviceIdl" : "../idl/teacher.thrift",
      	"serviceIdlContent" : "idl2"
      }
      ```

  * 查：\<IP address>/get-service

    * Body例：

      ```json
      { "serviceName" : "teacher-server"}
      ```

  * 查询所有IDL：\<IP address>/list-service



### 3.2. 扩展功能

* 负载均衡

  * 默认采用随机权重策略

  * 可在biz/clientprovider自定义负载均衡策略

  * ```go
    // clientprovider.go
    cli, err = genericclient.NewClient(serviceName, g, client.WithResolver(r),
                                       client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()))
    ```

* 路由选择

  * 默认提供 \<IP address>/gateway/ServiceName/MethodName 的网关接口

  * 可在router.go新增路由解析策略

  * ```go
    // router.go
    func customizedRegister(r *server.Hertz) {
    	r.GET("/ping", handler.Ping)
    
    	// your code ...
    	r.POST("/gateway/:service/:method", handler.CallServiceMethod)
    
    	routeInfo := r.Routes()
    	hlog.Info(routeInfo)
    }
    ```

    

* 跨域

  * 默认实现对任意origin的跨域，可按需求在main.go更改

  * ```go
    // main.go
    h := server.Default()
    
    h.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PUT", "PATCH"},
        AllowHeaders:     []string{"Origin", "Content-Type", "X-Auth-Token"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
    
    register(h)
    ```

* 服务发现

  * 默认使用etcd注册中心。可自定义使用其他服务注册中心

  * ```go
    // clientprovider.go
    r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
        if err != nil {
        panic(err)
    }
    ```

* 缓存时长

  * 默认配置30min的kitex client和 IDL对象缓存，可在clientprovider与IDL-provider自定义缓存时长

  * ```go
    // clientprovider.go
    func GetCacheClient(serviceName string) (genericclient.Client, error) {
    	if newClient, ok := clientCache.Load(serviceName); ok {
    		oldClient := newClient.(genericclient.Client)
    		cacheTimeout := 30 * time.Minute
    	}
    	...
    }
    ```

    



## 4.测试计划

### 4.1. 测试方案

两个功能点需要进行测试

* api gateway功能（核心功能）
  * router、kitex client provider、idl provider模块都与性能密切相关
  * 需测试接口
    * **localhost:8888/gateway/{ServiceName}/{MethodName}**
    * 以 localhost:8888/gateway/teacher-server/Query 为例
  * RPC服务端不是测试的重点，直接返回发送的原json体即可
* idl manager 模块（辅助功能）
  * 重点测试查询接口，因其最常用
  * 需测试接口
    * **localhost:8888/get-service**
    * localhost:8888/add-service
    * localhost:8888/update-service



### 4.2. 测试环境

测试工具

* Apache JMeter 5.6.2

测试配置

* 10 个线程循环运行300次/10000次
  * ![image-20230724231517730](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230724231517730.png)
  * ![image-20230724231502338](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230724231502338.png)
* 100 个线程循环运行100次/10000次
  * ![image-20230724224026610](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230724224026610.png)
  * ![image-20230724231450300](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230724231450300.png)
* 以.jmx文件保存

测试使用命令：

* 进入JMeter目录
  * cd .\apache-jmeter-5.6.2\bin
* 生成html测试报告
  * .\jmeter -n -t ..\jmx\cloudwego.jmx -l ..\results\log -e -o ..\results\output

### 4.3. 性能测试数据：

#### 10 个线程，循环运行300次

![image-20230724170702326](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230724170702326.png)



#### 100 个线程，循环运行100次

![image-20230724180117775](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230724180117775.png)

#### 性能分析结果

* 网关路由转发核心功能
  * 10个线程下，/gateway/{ServiceName}/{MethodName}接口平均RT响应时间为**6.76ms**，TPS吞吐量为**1006.71**
  * 100个线程下，/gateway/{ServiceName}/{MethodName}接口平均RT响应时间为**15.73ms**，TPS吞吐量为**804.89**
* IDL管理模块
  * 新增服务对应IDL偶尔出现死锁状况，但**管理平台对并发量要求不高**，可延迟优化
  * 查询服务对应IDL正常。接口RT响应时间小于1ms



## 5.优化计划

### 5.1. 可行性思考

优化思路：

* 算法和数据结构
  * 手动/用静态分析工具计算各循环复杂度
* 并发&并行
  * 另一线程自动更新IDL
* 网络和磁盘IO
  * 框架netpoll已做优化
  * 复用已经建立的链接，尽量保持当前连接不断开
* 内存分配&垃圾回收
  * 网关复用kitex client、IDL对象
  * kitex自带的对象池缓存
  * Go 自带的垃圾回收、逃逸分析
* 编译器优化 - 内联（inline）
  * 默认开启



优化分析：Go语言自带的 profiling 工具 

* pprof：分析程序运行时的 CPU、内存、阻塞等资源使用情况

* 火焰图：通过测量程序运行时各部分的相对执行时间，发现程序的热点，以便重点优化。

* 增量对比：抓取两次 pprof 数据，通过比对定位问题

* 观察线程：通过观察 goroutine stack 定位问题

  

### 5.2. 具体优化方案

通过分析过后，开发人员确定了性能瓶颈所在，并制定了具体优化方案

* kitex client provider复用kitex client缓存对象
  * 缓存过期时间为30min

```go
var clientCache sync.Map
var cacheTime = make(map[genericclient.Client]time.Time)

// GetCacheClient 缓存泛化调用客户端
func GetCacheClient(serviceName string) (genericclient.Client, error) {
	if newClient, ok := clientCache.Load(serviceName); ok {
		oldClient := newClient.(genericclient.Client)

		cacheTimeout := 30 * time.Minute

		if time.Since(cacheTime[oldClient]) < cacheTimeout {
			// 缓存未过期，直接返回
			return oldClient, nil
		}
		// 缓存已过期，删除旧缓存
		clientCache.Delete(serviceName)
	}

	// 缓存不存在，创建新缓存
	newClient, err := InitGenericClient(serviceName)
	if err != nil {
		return nil, err
	}
	cacheTime[newClient] = time.Now() // 记录缓存时间
	clientCache.Store(serviceName, newClient)
	return newClient, nil
}

```

* IDL provider复用IDL缓存对象
  * 缓存过期时间为30min

```go
var idlCache sync.Map
var cacheTime = make(map[*generic.ThriftContentProvider]time.Time)

// GetCacheIdl 缓存解析后的IDL内容
func GetCacheIdl(serviceName string) (*generic.ThriftContentProvider, error) {
	if idl, ok := idlCache.Load(serviceName); ok {
		oldIdl := idl.(*generic.ThriftContentProvider)
		cacheTimeout := 30 * time.Minute

		if time.Since(cacheTime[oldIdl]) < cacheTimeout {
			// 缓存未过期，直接返回
			return oldIdl, nil
		}
		// 缓存已过期，删除旧缓存
		idlCache.Delete(serviceName)
	}

	// 缓存不存在，创建新缓存
	idl, err := GetResolvedIdl(serviceName)
	if err != nil {
		return nil, err
	}
	cacheTime[idl] = time.Now() // 记录缓存时间
	idlCache.Store(serviceName, idl)

	return idl, nil
}

```

* idl manager 加锁防止map竞态

```go
func AddService(service gateway.Service) {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	if _, ok := ServiceNameMap[service.ServiceName]; !ok {
		ServiceNameMap[service.ServiceName] = service
	}
}
```

* idl manager 减少IO消耗
  * 在idl cache处已做优化



### 5.3. 优化后性能测试数据

#### 10 个线程，循环运行10000次

![image-20230724225330245](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230724225330245.png)

#### 100 个线程，循环运行10000次

![image-20230724225345015](https://lar-blog.oss-cn-nanjing.aliyuncs.com/picGo_img/AppData/Roaming/Typora/typora-user-images/image-20230724225345015.png)

#### 性能分析结果

* 网关路由转发核心功能
  * 10个线程下，/gateway/{ServiceName}/{MethodName}接口平均RT响应时间为**0.86ms**，TPS吞吐量为**7275.37**
    * 与6.76ms相比有显著提升
  * 100个线程下，/gateway/{ServiceName}/{MethodName}接口平均RT响应时间为**12.71ms**，TPS吞吐量为**6795.37**
    * 与15.73ms相比有较高提升，但仍响应较慢
    * 猜测是由于RPC服务端的性能限制
