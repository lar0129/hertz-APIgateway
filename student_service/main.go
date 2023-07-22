package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"kitex.demo/kitex_gen/demo/studentservice"
)

func main() {
	// start generic server
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	handler := new(StudentServiceImpl)
	svr := studentservice.NewServer(handler, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "student-server",
		Tags:        map[string]string{"Cluster": "student"},
	}))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
