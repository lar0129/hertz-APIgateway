package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"service.teacher/kitex_gen/teacher/teacherservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":9997")
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	handler := new(TeacherServiceImpl)
	svr := teacherservice.NewServer(handler, server.WithRegistry(r), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "teacher-server",
		Tags:        map[string]string{"Cluster": "teacher"},
	}))
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
