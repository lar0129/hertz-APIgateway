// Code generated by hertz generator.

package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/client/genericclient"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	clientprovider "hertz.demo/biz/clientprovider"
	demo "hertz.demo/biz/model/demo"
)

// Register .
// @router /add-student-info [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.Student
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("jsonReq:", string(jsonReq))

	// 泛化调用
	cli := clientprovider.GetGenericClient("student-server").(genericclient.Client)
	resp, err := cli.GenericCall(ctx, "Register", string(jsonReq))
	// resp := new(demo.RegisterResp)

	c.JSON(consts.StatusOK, resp)
}

// Query .
// @router /query [GET]
func Query(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.QueryReq
	err = c.BindAndValidate(&req) // c和req数据绑定
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	jsonReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("jsonReq:", string(jsonReq))

	// 泛化调用
	cli := clientprovider.GetGenericClient("student-server").(genericclient.Client)
	resp, err := cli.GenericCall(ctx, "Query", string(jsonReq))

	c.JSON(consts.StatusOK, resp)
}
