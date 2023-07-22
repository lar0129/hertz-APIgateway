package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"log"
	"net"

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

// type GenericServiceImpl struct {
// }

// func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
// 	// use jsoniter or other json parse sdk to assert request
// 	m := request.(string)
// 	fmt.Printf("Recv body: %v\n", m)
// 	fmt.Printf("Recv method: %v\n", method)
// 	studentServiceImpl := new(StudentServiceImpl)
// 	switch method {
// 	case "Query":
// 		var queryReq demo.QueryReq
// 		err = json.Unmarshal([]byte(m), &queryReq)
// 		if err != nil {
// 			fmt.Println("json parse error!", err)
// 			return
// 		}
// 		fmt.Println("QueryReq:", queryReq)
// 		response, err = studentServiceImpl.Query(ctx, &queryReq)
// 		fmt.Println("response:", response)
// 		return "success", nil
// 	case "Register":
// 		var registerReq demo.Student
// 		err = json.Unmarshal([]byte(m), &registerReq)
// 		if err != nil {
// 			fmt.Println("json parse error!", err)
// 			return
// 		}
// 		fmt.Println("registerReq:", registerReq)
// 		response, err = studentServiceImpl.Register(ctx, &registerReq)
// 		fmt.Println("response:", response)
// 		return "success", nil
// 	}

// 	return nil, err
// }
