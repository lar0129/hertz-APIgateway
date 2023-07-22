namespace go teacher

struct Teacher {
    1: required string name(api.body='name'),
}

struct QueryReq {
    1: required string name(api.body='name')
}

//----------------------service-------------------
service TeacherService {
    Teacher Query(1: QueryReq req)(api.post = '/get-teacher-info')
}