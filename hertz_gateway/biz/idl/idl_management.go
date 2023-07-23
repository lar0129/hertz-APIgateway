package idl

import (
	"hertz.demo/biz/model/gateway"
)

// idl management platform - backend
var serviceNameMap = make(map[string]gateway.Service)

func GetService(serviceName string) *gateway.Service {
	service := serviceNameMap[serviceName]
	return &service
}

func AddService(service gateway.Service) {
	if _, ok := serviceNameMap[service.ServiceName]; !ok {
		serviceNameMap[service.ServiceName] = service
	}
}

func DeleteService(serviceName string) {
	delete(serviceNameMap, serviceName)
}

func UpdateService(service gateway.Service) {
	if _, ok := serviceNameMap[service.ServiceName]; ok {
		serviceNameMap[service.ServiceName] = service
	}
}

func GetAllService() *[]*gateway.Service {
	var services []*gateway.Service
	for _, v := range serviceNameMap {
		services = append(services, &v)
	}
	return &services
}
