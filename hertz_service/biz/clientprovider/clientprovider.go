package clientprovider

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	etcd "github.com/kitex-contrib/registry-etcd"

	"hertz.demo/biz/idl"
)

var cli genericclient.Client

// 抽象出来的泛化调用方法
func GetGenericClient(ctx *context.Context, c *app.RequestContext) (response interface{}) {

	// 解析IDL文件
	// 直接解析
	// p, err := generic.NewThriftFileProvider("idl/stu.thrift")
	// if err != nil {
	// 	panic(err)
	// }

	// 动态解析
	p, err := idl.GetResolvedIdl()
	if err != nil {
		panic(err)
	}

	// 构造 JSON 请求和返回类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})

	// 负载均衡
	//client.WithTag("Cluster", "student")
	//client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer())

	// get the client
	//cli, err := genericclient.NewClient("student-server", g, client.WithHostPorts("127.0.0.1:9999"))
	cli, err = genericclient.NewClient("student-server", g, client.WithResolver(r))
	if err != nil {
		panic(err)
	}

	// resp is a client
	return cli
}
