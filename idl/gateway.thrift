namespace go gateway

//--------------------request & response--------------
struct Service {
    1: required string serviceName(go.tag = 'json:"serviceName"'),
    2: required string serviceIdlName(go.tag = 'json:"serviceIdl"'),
    3: string serviceIdlContent(go.tag = 'json:"serviceIdlContent"'),
}