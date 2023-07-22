package main

import (
	"context"

	demo "kitex.demo/kitex_gen/demo"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct{}

var studentsMap = make(map[int32]*demo.Student)

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...

	if _, ok := studentsMap[student.Id]; !ok {
		studentsMap[student.Id] = student
		resp = &demo.RegisterResp{
			Success: true,
			Message: "Regist success",
		}
	} else {
		resp = &demo.RegisterResp{
			Success: false,
			Message: "Regist failed, student already exists",
		}
	}

	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...
	resp = studentsMap[req.Id]
	if resp == nil {
		resp = &demo.Student{
			Id:      req.Id,
			Name:    "Not Found",
			College: &demo.College{},
			Email:   []string{"Not Found"},
		}
	}

	return
}
