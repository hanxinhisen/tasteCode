// Created by Hisen at 2019-07-11.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func c() {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	defer close(ch1)
}

func w() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go m()
	}
	go func() {
		wg.Wait()
		close(ch2)
	}()

}

func m() {
	for i := range ch1 {
		ch2 <- i
	}
	wg.Done()

}

var ch1 = make(chan int, 10)
var ch2 = make(chan int, 10)

func main() {
	go c()
	w()
	for i := range ch2 {
		fmt.Println(i)
	}
}
