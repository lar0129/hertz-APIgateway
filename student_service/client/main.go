package main

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/klog"
	"kitex.demo/kitex_gen/demo"
	"kitex.demo/kitex_gen/demo/studentservice"
)

func main() {
	var opts []client.Option
	opts = append(opts, client.WithHostPorts("127.0.0.1:9999"))
	opts = append(opts, client.WithLongConnection(connpool.IdleConfig{MinIdlePerAddress: 10,
		MaxIdlePerAddress: 1000}))
	stuCli := studentservice.MustNewClient("bytedance.kitex.demo", opts...)

	// r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	// stuCli := studentservice.MustNewClient("kitex.demo", client.WithResolver(r))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	req := &demo.QueryReq{
		Id: 1,
	}

	for i := 0; i < 3; i++ {
		resp, err := stuCli.Query(context.Background(), req)
		if err != nil {
			klog.Infof("SayHello failed: %s", err.Error())
		} else {
			klog.Infof("SayHello success: %s", resp)
		}
	}
}
