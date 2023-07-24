// idl management platform - backend

package idl

import (
	"errors"
	"hertz.demo/biz/model/gateway"
	"os"
	"sync"
)

var ServiceNameMap = make(map[string]gateway.Service)

var mapMutex = &sync.Mutex{}
var ioMutex = &sync.Mutex{}

func GetService(serviceName string) *gateway.Service {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	service, ok := ServiceNameMap[serviceName]
	if !ok {
		err := errors.New("service not found")
		panic(err)
	}
	return &service
}

func GetIdlPath(serviceName string) string {
	service := GetService(serviceName)
	return service.ServiceIdlName
}

// GetIdlContent 涉及到文件读写，需要加锁
func GetIdlContent(serviceName string) string {
	service := GetService(serviceName)
	ioMutex.Lock()
	defer ioMutex.Unlock()

	file, err := os.Open(service.ServiceIdlName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := os.ReadFile(service.ServiceIdlName)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func AddService(service gateway.Service) {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	if _, ok := ServiceNameMap[service.ServiceName]; !ok {
		ServiceNameMap[service.ServiceName] = service
	}
}

func DeleteService(serviceName string) {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	delete(ServiceNameMap, serviceName)
}

func UpdateService(service gateway.Service) {
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	if _, ok := ServiceNameMap[service.ServiceName]; ok {
		ServiceNameMap[service.ServiceName] = service
	}
}

func GetAllService() []*gateway.Service {
	var services []*gateway.Service
	mapMutex.Lock() // 获取锁定
	defer mapMutex.Unlock()
	for k := range ServiceNameMap {
		service := ServiceNameMap[k]
		services = append(services, &service)
	}
	return services
}
