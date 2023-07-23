package clientprovider

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"

	"hertz.demo/biz/idl"
)

var cli genericclient.Client

// 抽象出来的泛化调用方法
func GetGenericClient(serviceName string) (response interface{}) {

	// 解析IDL文件

	// 直接解析
	// p, err := generic.NewThriftFileProvider("idl/stu.thrift")
	// if err != nil {
	// 	panic(err)
	// }

	// 动态解析
	p, err := idl.GetResolvedIdl(serviceName)
	if err != nil {
		panic(err)
	}

	// 构造 JSON 请求和返回类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	// 在 etcd 中服务发现
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	// get the client from etcd
	//cli, err := genericclient.NewClient("student-server", g, client.WithHostPorts("127.0.0.1:9999")) // 直接连接
	cli, err = genericclient.NewClient(serviceName, g, client.WithResolver(r),
		client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()))
	if err != nil {
		panic(err)
	}

	// resp is a client
	return cli
}
