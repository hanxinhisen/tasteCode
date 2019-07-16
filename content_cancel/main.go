// Created by Hisen at 2019-06-19.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func worker(ctx context.Context) {

	for {
		fmt.Println("do work")
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			goto END
		default:

		}
	}
END:
	wg.Done()
}
func worker2(ctx context.Context) {

	for {
		fmt.Println("do work 2")
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			goto END
		default:

		}
	}
END:
	wg.Done()
}
func worker3(ctx context.Context) {

	for {
		fmt.Println("do work 3")
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			goto END
		default:

		}
	}
END:
	wg.Done()
}
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	wg.Add(3)
	go worker(ctx)
	go worker2(ctx)
	go worker3(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	wg.Done()
}
