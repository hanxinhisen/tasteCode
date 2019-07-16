// Created by Hisen at 2019-06-19.
package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("%v", err)
	}

}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	checkError(err)
	defer cli.Close()
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//_, err = cli.Put(ctx, "/aaaa", "111")
	//cancel()
	//checkError(err)
	//
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//res, err := cli.Get(ctx, "/aaaa")
	//checkError(err)
	//
	//for _, e := range res.Kvs {
	//	fmt.Println(string(e.Key), string(e.Value))
	//
	//}
	rch := cli.Watch(context.Background(), "aaaa")
	for w := range rch {
		for _, ev := range w.Events {
			fmt.Println(string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}
}
