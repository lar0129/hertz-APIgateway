package idl

import (
	"hertz.demo/biz/model/gateway"
)

// idl management platform - backend
var serViceNameMap = make(map[string]gateway.Service)

func GetService(serviceName string) *gateway.Service {
	service := serViceNameMap[serviceName]
	return &service
}

func AddService(service gateway.Service) {
	serViceNameMap[service.ServiceName] = service
}

func DeleteService(serviceName string) {
	delete(serViceNameMap, serviceName)
}

func UpdateService(service gateway.Service) {
	serViceNameMap[service.ServiceName] = service
}

func GetAllService() []*gateway.Service {
	var services []*gateway.Service
	for _, v := range serViceNameMap {
		services = append(services, &v)
	}
	return services
}
