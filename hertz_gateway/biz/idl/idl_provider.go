package idl

import (
	"errors"
	"github.com/cloudwego/kitex/pkg/generic"
	"io/ioutil"
	"time"
)

func GetResolvedIdl(serviceName string) (*generic.ThriftContentProvider, error) {
	// 动态解析
	// 做一个IDL文件缓存？

	service, ok := ServiceNameMap[serviceName]
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
	} else {
		// dynamic update
		// 仅当第一次获取IDL时，才创建新线程用来更新
		go func() {
			err = p.UpdateIDL(string(content), includes)
			if err != nil {
				panic(err)
			}
			time.Sleep(30 * time.Second)
		}()
	}

	return p, err
}
