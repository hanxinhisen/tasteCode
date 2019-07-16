// Created by Hisen at 2019-07-15.
package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

func work() {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(options)
	lockKey := "lock4counter"
	countKey := "counter"
	resp := client.SetNX(lockKey, 1, time.Second*5)

	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result", lockSuccess)
		return
	}

	getResp := client.Get(countKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++

		resp := client.Set(countKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("current countValue", cntValue)
	// 释放锁
	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil || unlockSuccess > 0 {
		fmt.Println("释放成功")
	} else {
		fmt.Println("释放失败")
	}

}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
}
