package idl

import (
	"github.com/cloudwego/kitex/pkg/generic"
	"sync"
	"time"
)

var idlCache sync.Map
var cacheTime = make(map[*generic.ThriftContentProvider]time.Time)

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
			err = p.UpdateIDL(content, includes)
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
		cacheTimeout := 30 * time.Minute

		if time.Since(cacheTime[oldIdl]) < cacheTimeout {
			// 缓存未过期，直接返回
			return oldIdl, nil
		}
		// 缓存已过期，删除旧缓存
		idlCache.Delete(serviceName)
	}

	// 缓存不存在，创建新缓存
	idl, err := GetResolvedIdl(serviceName)
	if err != nil {
		return nil, err
	}
	cacheTime[idl] = time.Now() // 记录缓存时间
	idlCache.Store(serviceName, idl)

	return idl, nil
}
