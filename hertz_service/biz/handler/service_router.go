package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client/genericclient"
	"hertz.demo/biz/clientprovider"
)

func CallServiceMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interface{}
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	service := c.Param("service")
	method := c.Param("method")
	fmt.Println("Call " + service + "'s method: " + method)

	// 将请求req参数转换为 JSON
	jsonReq, err := json.Marshal(req)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("jsonReq:", string(jsonReq))

	// 泛化调用
	cli := clientprovider.GetGenericClient(service).(genericclient.Client)
	resp, err := cli.GenericCall(ctx, method, string(jsonReq))

	c.JSON(consts.StatusOK, resp)
}
