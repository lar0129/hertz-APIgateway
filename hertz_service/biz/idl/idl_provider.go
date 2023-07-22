package idl

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/generic"
	"io/ioutil"
)

func GetResolvedIdl(serviceName string) (*generic.ThriftContentProvider, error) {
	// 动态解析

	service := serviceNameMap[serviceName]
	path := service.ServiceIdlName
	content, err := ioutil.ReadFile(path)
	fmt.Println("path:", path)
	fmt.Println("content:", string(content))

	if err != nil {
		panic(err)
	}
	includes := map[string]string{
		path: string(content),
	}
	// fmt.Println("includes:", includes)

	p, err := generic.NewThriftContentProvider(string(content), includes)
	if err != nil {
		panic(err)
	}
	// dynamic update
	err = p.UpdateIDL(string(content), includes)

	return p, err
}
