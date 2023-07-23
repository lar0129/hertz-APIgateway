package idl

import (
	"hertz.demo/biz/model/gateway"
)

// idl management platform - backend
var ServiceNameMap = make(map[string]gateway.Service)

func GetService(serviceName string) *gateway.Service {
	service := ServiceNameMap[serviceName]
	return &service
}

func AddService(service gateway.Service) {
	if _, ok := ServiceNameMap[service.ServiceName]; !ok {
		ServiceNameMap[service.ServiceName] = service
	}
}

func DeleteService(serviceName string) {
	delete(ServiceNameMap, serviceName)
}

func UpdateService(service gateway.Service) {
	if _, ok := ServiceNameMap[service.ServiceName]; ok {
		ServiceNameMap[service.ServiceName] = service
	}
}

func GetAllService() []*gateway.Service {
	var services []*gateway.Service
	for k := range ServiceNameMap {
		service := ServiceNameMap[k]
		services = append(services, &service)
	}
	return services
}
