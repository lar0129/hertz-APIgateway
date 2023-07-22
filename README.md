# hertz-APIgateway

## 测试方法
1. 打开服务端

* 从根目录打开kitex_service文件夹。
* 运行命令
  * sh build.sh
  * sh output/bootstrap.sh

2. 打开api网关服务

* 从根目录打开hertz_service文件夹。
* 运行命令
  * go build .
* 运行生成的可执行文件hertz.demo

3. 用postman或curl命令进行接口测试