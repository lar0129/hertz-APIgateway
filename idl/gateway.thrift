namespace go gateway

//--------------------request & response--------------
struct Service {
    1: required string serviceName(go.tag = 'json:"serviceName"'),
    2: required string serviceIdlName(go.tag = 'json:"serviceIdl"'),
    3: string serviceIdlContent(go.tag = 'json:"serviceIdlContent"'),
}

struct SuccessResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct ServiceReq {
    1: string serviceName(api.body='serviceName'),
}   

service IdlService {
    SuccessResp AddService(1: Service service)(api.post = '/add-service')
    SuccessResp DeleteService(1: ServiceReq serviceReq)(api.post = '/delete-service')
    SuccessResp UpdateService(1: Service service)(api.post = '/update-service')
    Service GetService(1: ServiceReq serviceReq)(api.post = '/get-service')
    list<Service> ListService()(api.post = '/list-service')
}