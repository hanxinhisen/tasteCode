// Created by Hisen at 2019-07-04.
package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/pkg/errors"
	"time"
)

var Number int
var Result string

func main() {
	config := hystrix.CommandConfig{
		Timeout:                2000,
		MaxConcurrentRequests:  8,
		SleepWindow:            1,
		ErrorPercentThreshold:  50,
		RequestVolumeThreshold: 5,
	}

	hystrix.ConfigureCommand("test", config)

	cbs, _, _ := hystrix.GetCircuit("test")
	defer hystrix.Flush()
	for i := 0; i < 1000; i++ {
		start1 := time.Now()
		Number = i
		hystrix.Do("test", runHandle(i, run), getFallBack)
		fmt.Println("请求次数:", i+1, ";用时:", time.Since(start1), ";请求状态 :", Result, ";熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
	}

}

func runHandle(i int, f func(int) error) func() error {

	ff := func() error {
		return f(i)
	}
	return ff

}
func run(i int) error {
	fmt.Println(i)
	Result = "Running1"
	if Number%2 == 0 {
		return nil
	} else {
		return errors.New("请求出错")
	}
}

func getFallBack(err error) error {

	Result = "FallBack"
	return nil

}
