package idl

import (
	"errors"
	"github.com/cloudwego/kitex/pkg/generic"
	"io/ioutil"
)

func GetResolvedIdl(serviceName string) (*generic.ThriftContentProvider, error) {
	// 动态解析

	service, ok := serviceNameMap[serviceName]
	if !ok {
		err := errors.New("service not found")
		return nil, err
	}
	path := service.ServiceIdlName
	content, err := ioutil.ReadFile(path)

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
