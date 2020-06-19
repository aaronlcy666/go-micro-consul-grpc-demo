package model

import (
	"context"
	"fmt"
	greeter "grpctest/server/proto"
)

//声明一个结构体

type Say struct {
}

//实现我们再proto中声明的方法自动生成的接口 三个参数为固定顺序 context.Context  自动生成的入参 自动生成的回参  返回 error
func (s *Say) Hello(ctx context.Context, req *greeter.SayParam, rsp *greeter.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*greeter.Pair)
	rsp.Header["name"] = &greeter.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = greeter.RespType_DESCEND

	return nil
}
