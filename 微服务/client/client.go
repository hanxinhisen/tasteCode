// Created by Hisen at 2019-07-09.
package main

import (
	"context"
	"fmt"

	proto "aaa/微服务/m/proto"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
	service.Init()
	helloservice := proto.NewIDGenerateClient("hanxin", service.Client())
	res, err := helloservice.GenID(context.TODO(), &proto.Request{Number: "3"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Msg)
	}

}
