package model

import (
	"context"
	"fmt"
	greeter "grpctest/server/proto"

)

type Say struct {

}

func (s *Say)Hello(ctx context.Context, req *greeter.SayParam, rsp *greeter.SayResponse) (error) {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*greeter.Pair)
	rsp.Header["name"] = &greeter.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = greeter.RespType_DESCEND

	return nil
}