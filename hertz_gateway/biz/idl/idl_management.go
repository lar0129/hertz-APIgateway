package idl

import (
	"hertz.demo/biz/model/gateway"
	"sync"
)

// idl management platform - backend
var ServiceNameMap = make(map[string]gateway.Service)
var mutex = &sync.Mutex{}

func GetService(serviceName string) *gateway.Service {
	mutex.Lock() // 获取锁定
	defer mutex.Unlock()
	service := ServiceNameMap[serviceName]
	return &service
}

func AddService(service gateway.Service) {
	mutex.Lock() // 获取锁定
	defer mutex.Unlock()
	if _, ok := ServiceNameMap[service.ServiceName]; !ok {
		ServiceNameMap[service.ServiceName] = service
	}
}

func DeleteService(serviceName string) {
	mutex.Lock() // 获取锁定
	defer mutex.Unlock()
	delete(ServiceNameMap, serviceName)
}

func UpdateService(service gateway.Service) {
	mutex.Lock() // 获取锁定
	defer mutex.Unlock()
	if _, ok := ServiceNameMap[service.ServiceName]; ok {
		ServiceNameMap[service.ServiceName] = service
	}
}

func GetAllService() []*gateway.Service {
	var services []*gateway.Service
	mutex.Lock() // 获取锁定
	defer mutex.Unlock()
	for k := range ServiceNameMap {
		service := ServiceNameMap[k]
		services = append(services, &service)
	}
	return services
}
