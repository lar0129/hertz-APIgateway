package idl

import (
	"github.com/cloudwego/kitex/pkg/generic"
	"sync"
	"time"
)

var idlCache sync.Map

func GetResolvedIdl(serviceName string) (*generic.ThriftContentProvider, error) {
	// 动态解析
	// 做一个IDL文件缓存？

	path := GetIdlPath(serviceName)
	content := GetIdlContent(serviceName)
	includes := map[string]string{
		path: content,
	}
	// fmt.Println("includes:", includes)

	p, err := generic.NewThriftContentProvider(content, includes)
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

// GetCacheIdl 缓存解析后的IDL内容
func GetCacheIdl(serviceName string) (*generic.ThriftContentProvider, error) {
	if idl, ok := idlCache.Load(serviceName); ok {
		oldIdl := idl.(*generic.ThriftContentProvider)
		return oldIdl, nil
	}
	idl, err := GetResolvedIdl(serviceName)
	if err != nil {
		return nil, err
	}
	idlCache.Store(serviceName, idl)
	return idl, nil
}
