// Created by Hisen at 2019-06-19.
package main

import (
	proto "aaa/微服务/m/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/sony/sonyflake"
	"os"
	"strconv"
)

func getHostname() (string, error) {
	return os.Hostname()
}

type ID struct {
}

func genID() (id string, err error) {
	setting := sonyflake.Settings{}
	setting.MachineID = func() (u uint16, e error) {
		return 1, nil
	}
	snowFlake := sonyflake.NewSonyflake(setting)
	idint64, err := snowFlake.NextID()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(idint64)), nil
}

func (h *ID) GenID(ctx context.Context, req *proto.Request, res *proto.Response) error {
	id, err := genID()
	hostname, _ := getHostname()
	if err != nil {
		res.Msg = "0"
	} else {
		res.Msg = hostname + ":" + id
	}
	return nil
}
func main() {
	service := micro.NewService(micro.Name("hanxin"))
	service.Init()
	proto.RegisterIDGenerateHandler(service.Server(), new(ID))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
