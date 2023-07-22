package idl

import (
	"github.com/cloudwego/kitex/pkg/generic"
	"io/ioutil"
)

func GetResolvedIdl() (*generic.ThriftContentProvider, error) {
	// 动态解析
	path := "../idl/stu.thrift"
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
